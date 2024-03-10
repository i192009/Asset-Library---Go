package logic

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful/contants"
	"assetLibary/restful/entries"
	"assetLibary/restful/httpDto"
	"assetLibary/services"
	"assetLibary/xutil"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.zixel.cn/go/framework/compoment"
	frameworkxUtil "gitlab.zixel.cn/go/framework/xutil"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var storage = compoment.NewOBS(configs.ObsAccessKey, configs.ObsSecretKey, configs.ObsEndPoint, configs.ObsDlBucket)

func AddAsset(appId string, entry *entries.CreateAssetEntry) (*entries.AddAssetRpn, error) {

	if ok, erron := CheckPermission(appId, entry.User, entry.Owner, entry.PermissionType, contants.OperateType_Write); !ok {
		return nil, erron
	}
	summary := map[string]interface{}{
		"appId":    appId,
		"filename": entry.FileName,
		"filesize": entry.FileSize,
		"type":     entry.Type,
		//"subtype":  entry.SubType,
		"platform": entry.Platform,
	}
	isJsonType := false
	var (
		assetUrl string
		path     string
		assetKey string
	)
	if entry.FileName == "" {
		isJsonType = true
	}
	if !isJsonType {
		//校验资产库配置
		if ok, err := CheckConfig(appId, entry.PermissionType, entry.Owner, entry.FileSize, 1); err != nil {
			return nil, err //ch
		} else if !ok {
			log.Error(err.Error())
			return nil, err //ch
		}
		if entry.Type == "" {
			fileExt := filepath.Ext(entry.FileName)
			if typeV, ok := frameworkxUtil.GetMimeTypes()[fileExt]; ok {
				entry.Type = typeV
			} else {
				log.Errorf("MIME Type is Failed")
				return nil, CreateError(codes.InvalidArgument, 50112, map[string]interface{}{
					"error": "MIME Type is Failed",
				})
			}
		}
		//校验分类配置
		fileExt := filepath.Ext(entry.FileName)
		if ok, err := CheckClass(entry.Class, entry.Type, fileExt, entry.FileSize); err != nil {

			return nil, err
		} else if !ok {
			log.Errorf("AddAsset>>>CheckClass  is error.err:%v", err.Error())
			return nil, err
		}
		path = frameworkxUtil.FileHashName(summary, []int{2, 2, 2, 4}, appId)
		if len(path) == 0 {
			log.Errorf("AddAsset>>>FileHashName  error.")
			return nil, CreateError(codes.Internal, 50113, map[string]interface{}{
				"error": "FileHashName error",
			})
		}

		assetKey = filepath.Join(configs.ObsUlFolder, path) + filepath.Ext(entry.FileName)

		UploadFileUrlRequest := &services.UploadFileUrlRequest{
			ScopeId: int32(configs.ObsScopeId),
			Key:     assetKey,
			Expires: time.Duration(300 + entry.ThumbnailSize/204800).Microseconds(),
			Size:    entry.FileSize,
		}

		if entry.Type != "" {
			UploadFileUrlRequest.ContentType = entry.Type
			UploadFileUrlRequest.CustomContentType = true
		}

		urlReply, err := services.StorageServiceService.GetUploadFileUrl(context.Background(), UploadFileUrlRequest)
		if err != nil {
			fmt.Println("Error", err)
			log.Errorf("AddAsset>>>GetUploadFileUrl is error.err:%v", err.Error())
			return nil, CreateError(codes.Internal, 50103, map[string]interface{}{
				"error": err.Error(),
			})
		}

		assetUrl = urlReply.Url
	}

	var thumbUrl string
	var thumbKey string
	if entry.Thumbnail != "" {
		if entry.ThumbnailSize == 0 {
			log.Errorf("AddAsset>>>thumbnail size or type empty", entry.ThumbnailSize, entry.ThumbnailType)
			//configs.SendBadReqeust(c, 48002, DefaultMessage)
			return nil, CreateError(codes.Internal, 50116, map[string]interface{}{
				"error": "thumbnail size or type empty",
			})
		}

		//thumbHeaders := map[string]string{
		//	"Content-Type":   "image/" + contants.ContentType_AssetSubType_Pic[entry.ThumbnailType],
		//	"Content-Length": strconv.FormatInt(entry.ThumbnailSize, 10),
		//}
		//
		//thumbUrl, err = Storage.GetUploadUrl(filepath.Join(configs.GetMyConfig().OBS_UL_Folder, path+".thumb"), 600, thumbHeaders)
		//if err != nil {
		//	xutil.Logger.Error(err.Error())
		//	//configs.SendInternalServerError(c, 1000, DefaultMessage)
		//	return
		//}
		if isJsonType {
			summaryTh := map[string]interface{}{
				"appId":    appId,
				"filename": entry.Thumbnail,
				"filesize": entry.ThumbnailSize,
				//"type":     entry.Type,
				//"subtype":  entry.SubType,
				"platform": entry.Platform,
			}
			pathTh := frameworkxUtil.FileHashName(summaryTh, []int{2, 2, 2, 4}, appId)
			thumbKey = filepath.Join(configs.ObsUlFolder, pathTh+".thumb")
		} else {
			thumbKey = filepath.Join(configs.ObsUlFolder, path+".thumb")
		}
		UploadFileUrlRequest := &services.UploadFileUrlRequest{
			ScopeId: int32(configs.ObsScopeId),
			Key:     thumbKey,
			Expires: time.Duration(300 + entry.FileSize/204800).Microseconds(),
			Size:    entry.ThumbnailSize,
		}
		if entry.ThumbnailType != "" {
			UploadFileUrlRequest.ContentType = entry.ThumbnailType
			UploadFileUrlRequest.CustomContentType = true
		}

		urlReply, err := services.StorageServiceService.GetUploadFileUrl(context.Background(), UploadFileUrlRequest)
		if err != nil {
			log.Errorf("AddAsset>>>GetUploadFileUrl  is error.err:%v", err.Error())
			return nil, CreateError(codes.Internal, 50103, map[string]interface{}{
				"error": err.Error(),
			})
		}
		thumbUrl = urlReply.Url
	}
	external := entry.External
	var otherFileMap map[string]string
	if entry.OtherFile != nil {
		if external == nil {
			external = make(map[string]interface{})
		}
		otherFileMap = make(map[string]string)
		for k, v := range entry.OtherFile {
			//headers := map[string]string{
			//	"Content-Type":    contants.ContentType_AssetType[v.Type] + "/" +  contants.ContentType[v.Type][v.SubType],
			//	"Content-Length": strconv.FormatInt(v.Size, 10),
			//}
			//url, err := services.StorageServiceService.GetUploadUrl(
			//	filepath.Join(configs.OBS_UL_Folder, path+"/"+k),
			//	time.Duration(300+entry.FileSize/204800)*time.Second,
			//	headers,
			//)
			UploadFileUrlRequest := &services.UploadFileUrlRequest{
				ScopeId: int32(configs.ObsScopeId),
				Key:     filepath.Join(configs.ObsUlFolder, path+"/"+k),
				Expires: time.Duration(300 + entry.FileSize/204800).Microseconds(),
				Size:    v.Size,
			}
			urlReply, err := services.StorageServiceService.GetUploadFileUrl(context.Background(), UploadFileUrlRequest)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			otherFileMap[k] = urlReply.Url
		}
		external["otherFile"] = otherFileMap
	}
	if entry.PermissionType == contants.AssetPermissionType_Private {
		entry.Owner = entry.User
	}

	asset := &dao.Asset{
		FileName: entry.FileName,
		FileSize: entry.FileSize,
		Type:     entry.Type,
		//SubType:        entry.SubType,
		Class: entry.Class,
		Tags:  entry.Tags,
		Title: entry.Title,
		AppId: appId,
		//InstanceId:     entry.InstanceId,
		Description: entry.Description,
		//Thumbnail:      entry.Thumbnail,
		Thumbnail:      thumbKey,
		Url:            assetKey,
		PermissionType: entry.PermissionType,
		Status:         contants.AssetStatus_New,
		Source:         entry.Source,
		External:       external,
		Creator:        entry.User,
		Owner:          entry.Owner,
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
		//RelatedAssets:  entry.RelatedAssets,
	}
	if isJsonType {
		asset.Status = contants.AssetStatus_Normal
	}
	C2S_UserInfoGetV2_Req := &services.C2S_UserInfoGetV2_Req{
		Id:           entry.User,
		InfoType:     1,
		InfoSelector: 0x4,
	}
	S2C_UserInfoGetV2_Rpn, err := services.UserServiceV2.GetUserInfo(context.Background(), C2S_UserInfoGetV2_Req)
	if err == nil && S2C_UserInfoGetV2_Rpn.PlainInfo != nil {
		asset.CreatorName = S2C_UserInfoGetV2_Rpn.PlainInfo.Name
	}
	if asset.Class != "" {
		if classDB, err := dao.GetClassById(entry.Class); err == nil {
			if classDB != nil {
				asset.ClassName = classDB.Name
			}
		}
	}
	if asset.Tags != nil && len(asset.Tags) > 0 {
		asset.TagsName = make([]string, len(asset.Tags))
		//Tags去除重复值
		asset.Tags = xutil.RemoveDuplicateElement(asset.Tags)
		if tages, err := dao.GetTagsByIds(asset.Tags); err == nil {
			if tages != nil {
				for i, tag := range tages {
					asset.TagsName[i] = tag.Name
				}
			}
		}
	}

	assetId, err := dao.AddAsset(asset)
	if err != nil {
		log.Errorf("AddAsset>>>dao.AddAsset  is error.err:%v", err.Error())
		return nil, CreateError(codes.Internal, 1001, map[string]interface{}{
			"error": err.Error(),
		}) //ch//&configs.ErrorNo{Code: 1001, Message: "dao.AddAsset is error"}
	}
	rpn := &entries.AddAssetRpn{
		AssetId:        assetId,
		AssetUploadUrl: assetUrl,
		ThumbUploadUrl: thumbUrl,
		OtherFile:      otherFileMap,
	}
	return rpn, nil
}

func AssetUploaded(assetId, permissionType string) error {
	asset, err := dao.GetAssetById(assetId, permissionType)
	if err != nil {
		return err
	}
	if asset == nil {
		log.Errorf("AssetUploaded>>> Asset is nil")
		return CreateError(codes.NotFound, 50118, map[string]interface{}{
			"error": "Asset not found",
		})
	}
	err = AdjustConfigSize(asset.Type, asset.Owner, asset.FileSize, 1)
	if err != nil {
		log.Errorf("AssetUploaded>>> AdjustConfigSize is error.err:%v", err.Error())
		return err
	}
	//var info xutil.RemoteFileInfo
	//err = xutil.GetRemoteFileinfo(
	//	fmt.Sprintf("https://%s.%s/%s",
	//		configs.OBS_UL_Bucket,
	//		configs.OBS_EndPoint,
	//		filepath.Join(configs.OBS_UL_Folder, asset.Thumbnail),
	//	),
	//	&info,
	//)
	isMustProc := false
	isThumb := false
	update := bson.M{"status": contants.AssetStatus_Normal}

	if asset.Thumbnail == "" {
		asset.Thumbnail = strings.TrimSuffix(asset.Url, filepath.Ext(asset.Url)) + ".thumb"
		isThumb = true
		update["thumbnail"] = asset.Thumbnail
	}
	if isMustProc || (isThumb && strings.Contains(asset.Type, contants.FileType_Pic)) {
		log.Infof("TransferFile init>>>>")
		//https://zetaverse.obs.cn-east-3.myhuaweicloud.com//Assets/9b/6f/94/a7f5efb60848bd23a095db47e0?x-image-process=image/resize,w_100,h_100/imageslim
		remoteUrl := fmt.Sprintf("https://%s.%s/%s?x-image-process=image/resize,l_%d/imageslim",
			configs.ObsDlBucket,
			configs.ObsEndPoint,
			asset.Url,
			400,
		)
		//DownLoadSignUrlRequest := &services.DownLoadSignUrlRequest{
		//	ScopeId: 4,
		//	Key:     asset.Url,
		//}
		//urlReply, err := services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
		//if err != nil {
		//	xutil.Logger.Errorf("GetAsset>>>GetUploadFileUrl  is error.err:%v", err)
		//	return errors.New("GetUploadFileUrl is error")
		//}
		////remoteUrl := fmt.Sprintf("%s&x-image-process=image/resize,l_%d/imageslim",
		////	urlReply.Url,
		////	400,
		////)
		//remoteUrl := urlReply.Url
		log.Infof("TransferFile init>>>>remoteUrl:%s,Thumbnail:%s", remoteUrl, asset.Thumbnail)
		err = storage.TransferFile(remoteUrl, asset.Thumbnail, nil)
		if err != nil {
			log.Error(err.Error())
		}
		asset.Thumbnail = remoteUrl
		//} else if isMustProc || (asset.Type == data.AssetType_File && asset.SubType == data.AssetSubType_File_3D && (len(asset.Thumbnail) == 0 || err != nil)) {
		//	if err := proc3D(asset.Assert_DB_Insert_t); err != nil {
		//		xutil.Logger.Error(err.Error())
		//	}

	} else if isMustProc || (isThumb && strings.Contains(asset.Type, contants.FileType_File)) {
		log.Infof("TransferFile proc3D>>>>")
		if err := proc3D(*asset); err != nil {
			log.Error(err.Error())
		}
	} else if isMustProc || (isThumb && strings.Contains(asset.Type, contants.FileType_Video)) {
		log.Infof("TransferFile procVideo>>>>")
		if err := procVideo(*asset); err != nil {
			log.Error(err.Error())
		}
	} else {
		log.Infof("TransferFile Thumbnail>>>>")
	}
	return dao.UpdateAssetById(assetId, asset.PermissionType, update)
}

func proc3D(asset dao.Asset_DB) error {
	// 请求realink 生成3d模型的缩略图
	var extern compoment.ConvertExtern_t
	extern.Marks = strings.Split(configs.RealinkMarks, ",")

	Job := compoment.NewConvertJob(configs.RealinkHttpAddr, configs.ObsEndPoint)
	jobId, err := Job.PostConvertJob(configs.ObsDlBucket,
		asset.Url,
		filepath.Join(configs.ObsUlFolder, filepath.Dir(asset.Url)),
		asset.FileName,
		[]string{"png"},
		&extern,
	)

	if err != nil {
		log.Errorf("PostConvertJob is error.err:%v", err.Error())
		log.Errorf("PostConvertJob is error.RealinkServiceAddr:%v,OBS_EndPoint:%v,", configs.RealinkHttpAddr, configs.ObsEndPoint)
		return CreateError(codes.InvalidArgument, 50120, map[string]interface{}{
			"error": err.Error(),
		})
	}
	log.Infof("PostConvertJob is ok.jobId:%v", jobId)

	go func(jobId string) {
		timer := time.NewTimer(time.Second * 5)
		point := time.Now().Add(time.Second*time.Duration(asset.FileSize/1024*1024) + time.Second*60)
		for {
			sig := <-timer.C

			timer.Reset(time.Second * 5)
			status, files, err := Job.GetConvertJobFiles(jobId)
			if err != nil {
				log.Error(err.Error())
				break
			}

			if status == 8 {
				log.Error("status = ", status)
				break
			}

			if point.Before(sig) {
				log.Error("wait time out")
				break
			}

			if status == 7 {
				log.Info("Convert completed. assetId = ", asset.AssetId, " convert file ", files[0])
				//OBS_UL_Folder:= MyConfig.OBS_UL_Folder
				//err = storage.CopyFile(files[0], filepath.Join(OBS_UL_Folder, asset.Thumbnail))
				//if err != nil {
				//	xutil.Logger.Error(err.Error())
				//	break
				//}
				copyObjectRequest := &services.CopyObjectRequest{
					ScopeId:   int32(configs.ObsScopeId),
					Key:       asset.Thumbnail,
					SourceKey: files[0],
					Bucket:    configs.ObsDlBucket,
				}
				log.Infof("CopyObjectRequest>>>>%v", copyObjectRequest)
				_, err = services.StorageServiceService.CopyObject(context.Background(), copyObjectRequest)
				if err != nil {
					log.Error(err.Error())
					break
				}

				err = storage.RemoveFile(files[0])
				if err != nil {
					log.Error(err.Error())
					break
				}

				break
			}
		}
	}(jobId)

	return nil
}

func procVideoObs(asset dao.Asset_DB) (err error) {
	var cwd string
	if cwd, err = os.Getwd(); err != nil {
		return errors.New("get current working directory error")
	}

	// / OBS object
	obs := compoment.NewOBS(configs.ObsDlBucket, configs.ObsSecretKey, configs.ObsEndPoint, configs.ObsUlBucket)

	videoUrl := filepath.Join(configs.ObsUlFolder, asset.Url)

	videoFile := asset.FileName
	videoFileExt := filepath.Ext(videoFile)
	videoFileName := strings.TrimRight(videoFile, videoFileExt)

	randDir := frameworkxUtil.RandomString(8)
	videoSaveTo := filepath.Join(cwd, "assets", randDir, videoFileName+".swp"+videoFileExt)

	if _, err = os.Stat(videoSaveTo); err != nil {
		err = obs.DownloadFile(videoUrl, videoSaveTo)
		if err != nil {
			return err
		}
	}

	imageUrl := filepath.Join(configs.ObsUlFolder, asset.Thumbnail)
	imageFile := filepath.Base(imageUrl)
	imageFileExt := filepath.Ext(imageUrl)
	imageSaveTo := filepath.Join(cwd, "assets", randDir, imageFile)

	err = xutil.VideoGenSnapshot(videoSaveTo, imageSaveTo, 1)
	if err != nil {
		return err
	}

	if err = obs.UploadFile(imageSaveTo, imageUrl, mime.TypeByExtension(imageFileExt)); err != nil {
		return err
	}

	go func() error {
		/// 这里可以继续优化，如果文件没有改变可以不用上传。
		defer os.RemoveAll(filepath.Join(cwd, "assets", randDir))

		if err = xutil.VideoMovFaststart(videoSaveTo, videoFile); err != nil {
			return err
		}

		if err = obs.UploadFile(videoFile, videoUrl, mime.TypeByExtension(videoFileExt)); err != nil {
			return err
		}

		return nil
	}()

	return nil
}

func procVideo(asset dao.Asset_DB) (err error) {
	var cwd string
	if cwd, err = os.Getwd(); err != nil {
		log.Error("Get current working directory error")
		return CreateError(codes.Internal, 50121, map[string]interface{}{
			"error": err.Error(),
		})
	}

	// / OBS object
	//obs := xutil.NewOBS(configs.OBS_AccessKey, configs.OBS_SecretKey, configs.OBS_EndPoint, configs.OBS_UL_Bucket)

	//videoUrl := filepath.Join(configs.OBS_UL_Folder, asset.Url)
	DownLoadSignUrlRequest := &services.DownLoadSignUrlRequest{
		ScopeId: int32(configs.ObsScopeId),
		Key:     asset.Url,
	}
	urlReply, err := services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
	if err != nil {
		log.Errorf("GetAsset>>>GetDownLoadSignUrl error.err:%v", err)
		return CreateError(codes.Internal, 50122, map[string]interface{}{
			"error": err.Error(),
		})
	}
	videoUrl := urlReply.Url
	videoFile := asset.FileName
	videoFileExt := filepath.Ext(videoFile)
	videoFileName := strings.TrimRight(videoFile, videoFileExt)

	randDir := frameworkxUtil.RandomString(8)
	videoSaveTo := filepath.Join(cwd, "assets", randDir, videoFileName+".swp"+videoFileExt)

	if _, err = os.Stat(videoSaveTo); err != nil {
		err = xutil.DownloadFile(videoUrl, videoSaveTo)
		//err = obs.DownloadFile(videoUrl, videoSaveTo)
		if err != nil {
			log.Errorf("DownloadFile error.err:%v", err)
			return CreateError(codes.InvalidArgument, 50123, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}

	//imageUrl := filepath.Join(configs.GetMyConfig().OBS_UL_Folder, asset.Thumbnail)
	//imageFile := filepath.Base(imageUrl)
	imageFile := filepath.Base(asset.Thumbnail)
	imageSaveTo := filepath.Join(cwd, "assets", randDir, imageFile)
	err = xutil.VideoGenSnapshot(videoSaveTo, imageSaveTo, 1)
	if err != nil {
		log.Errorf("VideoGenSnapshot is error.err:%v", err)
		return CreateError(codes.InvalidArgument, 50124, map[string]interface{}{
			"error": err.Error(),
		})
	}
	file, err := os.Open(imageSaveTo)
	if err != nil {
		log.Errorf("Open file error.err:%v", err)
		return CreateError(codes.InvalidArgument, 50125, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Errorf("File.Stat error.err:%v", err)
		return CreateError(codes.InvalidArgument, 50125, map[string]interface{}{
			"error": err.Error(),
		})
	}
	UploadFileUrlRequest := &services.UploadFileUrlRequest{
		ScopeId: int32(configs.ObsScopeId),
		Key:     asset.Thumbnail,
		Expires: 300 + fileInfo.Size()/204800,
		Size:    fileInfo.Size(),
	}
	urlReply, err = services.StorageServiceService.GetUploadFileUrl(context.Background(), UploadFileUrlRequest)
	if err != nil {
		log.Errorf("GetAsset>>>GetUploadFileUrl  is error.err:%v", err)
		return CreateError(codes.InvalidArgument, 50103, map[string]interface{}{
			"error": err.Error(),
		})
	}
	if err = frameworkxUtil.UploadFile(urlReply.Url, imageSaveTo); err != nil {
		//if err = obs.UploadFile(imageSaveTo, imageUrl, mime.TypeByExtension(imageFileExt)); err != nil {
		log.Errorf("Upload file error")
		return CreateError(codes.InvalidArgument, 50127, map[string]interface{}{
			"error": err.Error(),
		})
	}

	go func() error {
		/// 这里可以继续优化，如果文件没有改变可以不用上传。
		defer os.RemoveAll(filepath.Join(cwd, "assets", randDir))

		return nil
	}()

	return nil
}

func DeleteAsset(userId, appId, assetId, permissionType string) error {
	asset, err := dao.GetAssetById(assetId, permissionType)
	if err != nil {
		if err != nil {
			return err //&configs.ErrorNo{Err: err}
		}
	}
	if ok, erron := CheckPermission(appId, userId, asset.Owner, permissionType, contants.OperateType_Delete); !ok {
		return erron
	}
	if errn := dao.DeleteAsset(assetId, asset.PermissionType); errn != nil {
		return errn //&configs.ErrorNo{Err: err}
	}
	return nil
}

func UpdateAsset(appId, userId, assetId, permissionType string, entry *entries.UpdateAssetEntry) error {
	asset, err := dao.GetAssetById(assetId, permissionType)
	if err != nil {
		return err //&configs.ErrorNo{Err: err}
	}
	if ok, erron := CheckPermission(appId, userId, asset.Owner, permissionType, contants.OperateType_Write); !ok {
		return erron
	}
	if entry.Description != "" {
		asset.Description = entry.Description
	}
	if entry.Tags != nil {
		if len(entry.Tags) > 0 {
			asset.Tags = entry.Tags
			asset.TagsName = make([]string, len(asset.Tags))
			//Tags去除重复值
			asset.Tags = xutil.RemoveDuplicateElement(asset.Tags)
			if tages, err := dao.GetTagsByIds(asset.Tags); err == nil {
				if tages != nil {
					for i, tag := range tages {
						asset.TagsName[i] = tag.Name
					}
				}
			}
		} else {
			asset.Tags = []string{}
			asset.TagsName = []string{}
		}
	}
	if entry.Title != "" {
		asset.Title = entry.Title
	}
	if entry.External != nil {
		asset.External = entry.External
	}
	if entry.Thumbnail != "" {
		asset.Thumbnail = entry.Thumbnail
	}
	asset.UpdateTime = time.Now()
	if erron := dao.UpdateAsset(assetId, &asset.Asset); erron != nil {
		return erron //&configs.ErrorNo{Err: err}
	}
	return nil
}

func UpdateAssetThumbnail(appId, userId, assetId, permissionType string, entry *entries.UpdateAssetThumbnailEntry) (string, error) {

	asset, err := dao.GetAssetById(assetId, permissionType)
	if err != nil || asset == nil {
		return "", err
	}
	UploadFileUrlRequest := &services.UploadFileUrlRequest{
		ScopeId: int32(configs.ObsScopeId),
		Key:     asset.Thumbnail,
		Expires: time.Duration(300 + entry.ThumbnailSize/204800).Microseconds(),
		Size:    entry.ThumbnailSize,
	}

	if entry.ThumbnailType != "" {
		UploadFileUrlRequest.ContentType = entry.ThumbnailType
		UploadFileUrlRequest.CustomContentType = true
	}

	urlReply, err := services.StorageServiceService.GetUploadFileUrl(context.Background(), UploadFileUrlRequest)
	if err != nil {
		log.Errorf("AddAsset>>>GetUploadUrl  is error.err:%v", err)
		return "", CreateError(codes.InvalidArgument, 50103, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return urlReply.Url, nil
}

func QueryAssets(appId, userId, permissionType string, query *entries.GetAssetListReq_t) (*httpDto.GetAssetListRpn_t, error) {

	//校验权限
	if ok, erron := CheckPermission(appId, userId, query.Owner, permissionType, contants.OperateType_Read); !ok {
		return nil, erron
	}
	//openId := c.Request.Header.Get(ZixelHeader_OpenId)
	//if !CheckIf(c, data.Role_Editor) {
	//	return
	//}

	//query := GetAssetListReq_t{
	//	Page: 0,
	//	Size: 10,
	//	Type: 0,
	//	Tags: "",
	//	//AlbumId: data.AssetAlbum,
	//}

	//if err := c.ShouldBindQuery(&query); err != nil {
	//	xutil.Logger.Error(err.Error())
	//	configs.SendBadReqeust(c, 40001, DefaultMessage)
	//	return
	//}
	query.PermissionType = permissionType
	if query.Size == 0 {
		query.Size = query.PageSize
	}
	if query.Size > 100 {
		query.Size = 100
	}
	if query.Page <= 1 {
		query.Page = 0
	} else {
		query.Page = query.Page - 1
	}

	filter := bson.M{
		//data.AssetField_User: c.Request.Header.Get(ZixelHeader_OpenId),
		//data.AssetField_AlbumId: query.AlbumId,
	}
	//filter[dao.AssetField_AppId] = appId
	filter[dao.AssetField_Status] = contants.AssetStatus_Normal
	if query.Type > 0 {
		filter[dao.AssetField_Type] = query.Type
	}

	if query.Tags != "" {
		tags := strings.Split(query.Tags, ",")
		filter[dao.AssetField_Tags] = bson.D{{Key: "$in", Value: tags}}
	}
	if query.Platform != "" {
		filter[dao.AssetField_Platform] = query.Platform
	}
	if query.Class != "" {
		filter[dao.AssetField_Class] = query.Class
	}
	if query.PermissionType != "" {
		filter[dao.AssetField_PermissionType] = query.PermissionType
		if query.PermissionType == contants.AssetPermissionType_Private {
			filter[dao.AssetField_Owner] = userId
		} else if query.PermissionType == contants.AssetPermissionType_Tenant {
			filter[dao.AssetField_Owner] = query.Owner
		} else {
			if query.Owner != "" {
				filter[dao.AssetField_Owner] = query.Owner
			}
		}
	} else {
		filter["$or"] = []bson.M{{dao.AssetField_PermissionType: contants.AssetPermissionType_Public}, {dao.AssetField_PermissionType: ""}, {"$and": []bson.M{{dao.AssetField_User: query.Owner}}}}
	}
	//根据Search模糊查询
	if query.Search != "" {
		filter["$or"] = []bson.M{{dao.AssetField_Title: bson.M{"$regex": query.Search, "$options": "$i"}}, {dao.AssetField_Description: bson.M{"$regex": query.Search, "$options": "$i"}}}
	}
	if query.Order == "" {
		query.Order = "updateTime"
	}
	if query.Sort == 0 {
		query.Sort = -1
	}
	assetDBs, total, err := dao.QueryAssetByPage(permissionType, filter, query.Page, query.Size, query.Order, query.Sort)
	if err != nil {
		return nil, err
	}
	//totalCount, err := asset.CountDocuments(ctx, filter)
	//if err != nil {
	//	log.Error(err.Error())
	//	configs.SendInternalServerError(c, 1001, DefaultMessage)
	//	return
	//}
	result := make([]*httpDto.AssetDto, len(assetDBs))
	for i, assetDB := range assetDBs {
		assetDto := &httpDto.AssetDto{
			Id:             assetDB.AssetId,
			FileName:       assetDB.FileName,
			FileSize:       assetDB.FileSize,
			Type:           assetDB.Type,
			Class:          assetDB.Class,
			ClassName:      assetDB.ClassName,
			Tags:           assetDB.Tags,
			TagsName:       assetDB.TagsName,
			Thumbnail:      assetDB.Thumbnail,
			Title:          assetDB.Title,
			Description:    assetDB.Description,
			External:       assetDB.External,
			PermissionType: assetDB.PermissionType,
			Creator:        assetDB.Creator,
			CreatorName:    assetDB.CreatorName,
			//Platform:       assetDB.Platform,
			Owner:      assetDB.Owner,
			CreateTime: assetDB.CreateTime,
			UpdateTime: assetDB.UpdateTime,
			Status:     assetDB.Status,
		}
		DownLoadSignUrlRequest := &services.DownLoadSignUrlRequest{
			ScopeId: int32(configs.ObsScopeId),
			Key:     assetDto.Thumbnail,
		}
		urlReply, err := services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
		if err != nil {
			log.Infof("Url is missing. assetId = %s, thumbnail = %s", assetDto.Id, assetDto.Thumbnail)
		} else {
			assetDto.ThumbnailDownloadUrl = urlReply.Url
		}
		DownLoadSignUrlRequest.Key = assetDto.Url
		DownLoadSignUrlRequest.ScopeId = int32(configs.ObsScopeId)
		urlReply, err = services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
		if err != nil {
			log.Infof("Url is missing. assetId = %s, thumbnail = %s", assetDto.Id, assetDto.Thumbnail)
		} else {
			assetDto.AssetDownloadUrl = urlReply.Url
		}

		if assetDto.CreatorName == "" {
			C2S_UserInfoGetV2_Req := &services.C2S_UserInfoGetV2_Req{
				Id:           assetDto.Creator,
				InfoType:     1,
				InfoSelector: 0x4,
			}
			S2C_UserInfoGetV2_Rpn, err := services.UserServiceV2.GetUserInfo(context.Background(), C2S_UserInfoGetV2_Req)
			if err == nil {
				if S2C_UserInfoGetV2_Rpn != nil && S2C_UserInfoGetV2_Rpn.PlainInfo != nil {
					assetDto.CreatorName = S2C_UserInfoGetV2_Rpn.PlainInfo.Name
				}
			}
		}
		result[i] = assetDto
	}

	var rpn = &httpDto.GetAssetListRpn_t{
		Page:    query.Page,
		Size:    query.Size,
		Total:   total,
		Results: result,
	}
	return rpn, nil

}

func GetAsset(userId, appId, assetId, permissionType string) (*httpDto.AssetDto, error) {
	assetDB, err := dao.GetAssetById(assetId, permissionType)
	if err != nil || assetDB == nil {
		return nil, err //&configs.ErrorNo{Err: err}
	}
	assetDto := &httpDto.AssetDto{
		Id:             assetDB.AssetId,
		FileName:       assetDB.FileName,
		FileSize:       assetDB.FileSize,
		Type:           assetDB.Type,
		Class:          assetDB.Class,
		ClassName:      assetDB.ClassName,
		Tags:           assetDB.Tags,
		TagsName:       assetDB.TagsName,
		Title:          assetDB.Title,
		InstanceId:     assetDB.InstanceId,
		AppId:          assetDB.AppId,
		Description:    assetDB.Description,
		Thumbnail:      assetDB.Thumbnail,
		Url:            assetDB.Url,
		PermissionType: assetDB.PermissionType,
		Status:         assetDB.Status,
		Source:         assetDB.Source,
		External:       assetDB.External,
		Creator:        assetDB.Creator,
		CreatorName:    assetDB.CreatorName,
		Owner:          assetDB.Owner,
		RelatedAssets:  assetDB.RelatedAssets,
		CreateTime:     assetDB.CreateTime,
		UpdateTime:     assetDB.UpdateTime,
	}
	DownLoadSignUrlRequest := &services.DownLoadSignUrlRequest{
		ScopeId: int32(configs.ObsScopeId),
		Key:     assetDto.Thumbnail,
	}
	urlReply, err := services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
	if err != nil {
		log.Infof("Url is missing. assetId = %s, thumbnail = %s", assetDto.Id, assetDto.Thumbnail)
	} else {
		assetDto.ThumbnailDownloadUrl = urlReply.Url
	}

	log.Debugf("GetAsset>>>GetUploadFileUrl  is error.urlReply:%v", urlReply)

	DownLoadSignUrlRequest.Key = assetDto.Url
	DownLoadSignUrlRequest.ScopeId = int32(configs.ObsScopeId)
	urlReply, err = services.StorageServiceService.GetDownLoadSignUrl(context.Background(), DownLoadSignUrlRequest)
	if err != nil {
		log.Infof("Url is missing. assetId = %s, thumbnail = %s", assetDto.Id, assetDto.Thumbnail)
	} else {
		assetDto.AssetDownloadUrl = urlReply.Url
	}

	if assetDto.CreatorName == "" {
		C2S_UserInfoGetV2_Req := &services.C2S_UserInfoGetV2_Req{
			Id:           assetDto.Creator,
			InfoType:     1,
			InfoSelector: 0x4,
		}
		S2C_UserInfoGetV2_Rpn, err := services.UserServiceV2.GetUserInfo(context.Background(), C2S_UserInfoGetV2_Req)
		if err == nil && S2C_UserInfoGetV2_Rpn != nil && S2C_UserInfoGetV2_Rpn.PlainInfo != nil {
			assetDto.CreatorName = S2C_UserInfoGetV2_Rpn.PlainInfo.Name
		}
	}
	return assetDto, nil
}

func UseAsset(userId, appId, assetId string, permissionType string, copyReq *entries.CopyAssetReq_t) (string, error) {
	assetDB, err := dao.GetAssetById(assetId, permissionType)
	if err != nil {
		log.Errorf("UseAsset>>>GetAssetById  is error.err:%v", err)
		return "", err
	}
	if ok, erron := CheckPermission(appId, userId, assetDB.Owner, permissionType, contants.OperateType_Read); !ok {
		return "", erron //erron.Error()
	}

	ObsInstanceMap := configs.ObsInstanceMap
	targetPosition := configs.ObsDlBucket
	isHaveD := false //是否有D
	if InstanceScopeId_DB, err := GetInstanceScopeId(appId); err == nil {
		if InstanceScopeId_DB != nil {
			for k, v := range InstanceScopeId_DB.ScopeIdMap {
				if v == "D" {
					targetPosition = k
					isHaveD = true
					break
				}
			}
		}
	}
	if !isHaveD {
		if vaule, ok := ObsInstanceMap[copyReq.TargetPosition]; ok {
			targetPosition = vaule.(string)
		}
	}
	copyObjectRequest := &services.CopyObjectRequest{
		ScopeId:   int32(configs.ObsScopeId),
		Key:       assetDB.Url,
		SourceKey: assetDB.Url,
		Bucket:    targetPosition,
	}
	log.Infof("UseAsset>>>CopyObject  is copyObjectRequest:%v", copyObjectRequest)
	_, err = services.StorageServiceService.CopyObject(context.Background(), copyObjectRequest)
	if err != nil {
		log.Errorf("UseAsset>>>CopyObject error.err:%v", err)
		return "", CreateError(codes.Internal, 50129, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return assetDB.Url, nil
}

func GetPermissions(userId, appId string, checkPermissionReq_t *entries.CheckPermissionReq_t) (map[string]bool, error) {
	if checkPermissionReq_t.OperateType != "" {
		if ok, erron := CheckPermission(appId, userId, checkPermissionReq_t.TenantId, checkPermissionReq_t.PermissionType, checkPermissionReq_t.OperateType); !ok {
			log.Errorf("GetPermissions>>>CheckPermission is error.err:%v", erron)
			return map[string]bool{checkPermissionReq_t.OperateType: false}, erron
		} else {
			return map[string]bool{checkPermissionReq_t.OperateType: true}, nil
		}
	}
	if checkPermissionReq_t.PermissionType == contants.AssetPermissionType_Private {
		if checkPermissionReq_t.AssetId == "" {
			return map[string]bool{contants.OperateType_Write: true, contants.OperateType_Read: false, contants.OperateType_Delete: false}, nil
		} else {
			assetDB, err := dao.GetAssetById(checkPermissionReq_t.AssetId, checkPermissionReq_t.PermissionType)
			if err != nil || assetDB == nil {
				log.Errorf("GetPermissions>>>GetAssetById  is error.err:%v", err)
				return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: false, contants.OperateType_Delete: false}, err
			}
			if assetDB.Owner == userId {
				return map[string]bool{contants.OperateType_Write: true, contants.OperateType_Read: true, contants.OperateType_Delete: true}, nil
			} else {
				return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: false, contants.OperateType_Delete: false}, nil
			}
		}
	}
	if checkPermissionReq_t.PermissionType == contants.AssetPermissionType_Tenant {
		//TenantID为空，则无权限
		if checkPermissionReq_t.TenantId == "" {
			return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: false, contants.OperateType_Delete: false}, nil
		}
		//TenantID不为空
		C2S_ListUserByUIdReq := &services.C2S_ListUserByUIdReq{
			Uid:       []string{userId},
			CompanyId: checkPermissionReq_t.TenantId,
			AppId:     appId,
		}
		log.Infof("GetPermissions>>>ListUserByUid  is C2S_ListUserByUIdReq:%+v", C2S_ListUserByUIdReq)
		C2S_ListUserByOpenIdRpn, err := services.OrgMagService.ListUserByUid(context.Background(), C2S_ListUserByUIdReq)
		if err != nil {
			log.Errorf("GetPermissions>>>ListUserByUid  is error.err:%v", err)
			return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: false, contants.OperateType_Delete: false}, CreateError(codes.InvalidArgument, 50130, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//用户不存在，则无权限
		if len(C2S_ListUserByOpenIdRpn.UserInfo) == 0 {
			return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: false, contants.OperateType_Delete: false}, CreateError(codes.InvalidArgument, 50130, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//用户存在，判断角色
		if C2S_ListUserByOpenIdRpn.UserInfo[0].Role == "1" { // 管理员角色
			return map[string]bool{contants.OperateType_Write: true, contants.OperateType_Read: true, contants.OperateType_Delete: true}, nil
		}
		//非管理员角色
		return map[string]bool{contants.OperateType_Write: true, contants.OperateType_Read: true, contants.OperateType_Delete: false}, nil
	}
	if checkPermissionReq_t.PermissionType == contants.AssetPermissionType_Public {
		return map[string]bool{contants.OperateType_Write: false, contants.OperateType_Read: true, contants.OperateType_Delete: false}, nil
	}
	return nil, nil
}
