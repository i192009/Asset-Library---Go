package restful

import (
	"assetLibary/restful/entries"
	"assetLibary/restful/httpDto"
	"assetLibary/restful/logic"
	"assetLibary/restful/router"
	"assetLibary/services"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"gitlab.zixel.cn/go/framework/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"strings"
)

var log = logger.Get()

type AssetLibraryService struct {
	services.UnimplementedAssetLibrarySerivceServer
}

func NewService() *AssetLibraryService {
	return &AssetLibraryService{}
}
func (s *AssetLibraryService) QueryAssetPage(ctx context.Context, req *services.QueryAssetPageRequest) (*services.QueryAssetPageResponse, error) {
	defer func() {
		if r := recover(); r != nil {

			log.Errorf("panic: %v", r)
		}
	}()
	// TODO: 实现具体的查询逻辑
	query := &entries.GetAssetListReq_t{
		Page:           req.Page,
		Size:           req.PageSize,
		Type:           int(req.Type),
		Platform:       req.Platform,
		PermissionType: req.PermissionType,
		Owner:          req.Owner,
	}
	if req.Tags != nil && len(req.Tags) > 0 {
		query.Tags = strings.Join(req.Tags, ",")
	}
	if req.Class != nil {
		query.Class = *req.Class
	}
	if req.Order != nil {
		query.Order = *req.Order
	}
	if req.Sort != nil {
		query.Sort = int(*req.Sort)
	}
	if req.Search != nil {
		query.Search = *req.Search
	}
	GetAssetListRpn_t, errn := logic.QueryAssets(req.AppId, req.UserId, req.PermissionType, query)
	if errn != nil {
		errCode := SendErrorToJson(errn)
		return &services.QueryAssetPageResponse{
			Code:    errCode,
			Message: errn.Error(),
		}, errn
	}
	if GetAssetListRpn_t != nil {
		response := &services.QueryAssetPageResponse{
			Code:     200,
			Message:  "success",
			Total:    GetAssetListRpn_t.Total,
			Page:     GetAssetListRpn_t.Page,
			PageSize: GetAssetListRpn_t.Size,
		}
		if GetAssetListRpn_t.Results != nil && len(GetAssetListRpn_t.Results) > 0 {
			assets := make([]*services.Asset, len(GetAssetListRpn_t.Results))
			for index, assetDto := range GetAssetListRpn_t.Results {
				asset := &services.Asset{
					Id:                   assetDto.Id,
					FileName:             assetDto.FileName,
					FileSize:             assetDto.FileSize,
					Type:                 assetDto.Type,
					Class:                assetDto.Class,
					ClassName:            assetDto.ClassName,
					Tags:                 assetDto.Tags,
					TagsName:             assetDto.TagsName,
					Title:                assetDto.Title,
					InstanceId:           assetDto.InstanceId,
					AppId:                assetDto.AppId,
					Description:          assetDto.Description,
					Thumbnail:            assetDto.Thumbnail,
					ThumbnailDownloadUrl: assetDto.ThumbnailDownloadUrl,
					Url:                  assetDto.Url,
					AssetDownloadUrl:     assetDto.AssetDownloadUrl,
					PermissionType:       assetDto.PermissionType,
					Status:               int32(assetDto.Status),
					Source:               assetDto.Source,
					Creator:              assetDto.Creator,
					CreatorName:          assetDto.CreatorName,
					Owner:                assetDto.Owner,
					RelatedAssets:        assetDto.RelatedAssets,
					CreateTime:           assetDto.CreateTime.Format("2006-01-02 15:04:05"),
					UpdateTime:           assetDto.UpdateTime.Format("2006-01-02 15:04:05"),
				}
				if assetDto.External != nil && len(assetDto.External) > 0 {
					external, err := convertToStructMap(assetDto.External)
					if err != nil {
						asset.External = external
					}
				}
				assets[index] = asset
			}
			response.Asset = assets
		}

		return response, nil
	}
	// 返回一个 QueryAssetPageResponse 实例和 nil 错误
	return &services.QueryAssetPageResponse{}, nil
}

// 将 map[string]interface{} 转换为 map[string]*structpb.Struct
func convertToStructMap(src map[string]interface{}) (map[string]*structpb.Struct, error) {
	if src == nil {
		return nil, nil
	}
	result := make(map[string]*structpb.Struct)

	for k, v := range src {
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			log.Errorf("convertToStructMap error: %v", err)
			return nil, err
		}

		pbStruct := &structpb.Struct{}
		err = jsonpb.UnmarshalString(string(jsonBytes), pbStruct)
		if err != nil {
			log.Errorf("convertToStructMap error: %v", err)
			return nil, err
		}
		result[k] = pbStruct
	}
	return result, nil
}
func convertToMap(src map[string]*structpb.Struct) (map[string]interface{}, error) {
	if src == nil {
		return nil, nil
	}
	result := make(map[string]interface{})

	for k, v := range src {
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			log.Errorf("convertToMap error: %v", err)
			return nil, err
		}

		var iface interface{}
		err = json.Unmarshal(jsonBytes, &iface)
		if err != nil {
			log.Errorf("convertToMap error: %v", err)
			return nil, err
		}

		result[k] = iface
	}

	return result, nil
}

// 实现 AssetLibrarySerivceServer 接口中的 AddAsset 方法
func (s *AssetLibraryService) AddAsset(ctx context.Context, req *services.AddAssetRequest) (*services.AddAssetResponse, error) {
	// TODO: 实现具体的添加逻辑
	entry := &entries.CreateAssetEntry{
		FileName:       req.Name,
		FileSize:       req.Size,
		Type:           req.Type,
		Tags:           req.Tags,
		Class:          req.Class,
		Platforms:      req.Platforms,
		PermissionType: req.PermissionType,
		Title:          req.Title,
		AppId:          req.AppId,
		InstanceId:     req.InstanceId,
		Description:    req.Description,
		//External             :req.External      ,
		Thumbnail:     req.Thumbnail,
		ThumbnailSize: req.ThumbnailSize,
		ThumbnailType: req.ThumbnailType,
		//OtherFile            :req.OtherFile     ,
		User:   req.UserId,
		Owner:  req.Owner,
		Source: req.Source,
	}
	if external, err := convertToMap(req.External); err == nil {
		entry.External = external
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	AddAssetRpn, errn := logic.AddAsset(req.AppId, entry)
	if errn != nil {
		errCode := SendErrorToJson(errn)
		return &services.AddAssetResponse{
			Code:    errCode,
			Message: errn.Error(),
		}, errn
	}
	// 返回一个 AddAssetResponse 实例和 nil 错误
	return &services.AddAssetResponse{
		Code:           200,
		Message:        "success",
		AssetId:        AddAssetRpn.AssetId,
		AssetUploadUrl: AddAssetRpn.AssetUploadUrl,
		ThumbUploadUrl: AddAssetRpn.ThumbUploadUrl,
		OtherFiles:     AddAssetRpn.OtherFile,
	}, nil
}

/**
 * 实现 AssetLibrarySerivceServer 接口中的 AssetUploaded 方法
 */
func (s *AssetLibraryService) AssetUploaded(ctx context.Context, req *services.AssetUploadedRequest) (*services.AssetUploadedResponse, error) {
	if err := logic.AssetUploaded(req.Id, req.PermissionType); err != nil {
		errCode := SendErrorToJson(err)
		return &services.AssetUploadedResponse{
			Code:    errCode,
			Message: err.Error(), //configs.ErrorCodes[50106],
		}, err
	}
	return &services.AssetUploadedResponse{
		Code:    200,
		Message: "success",
	}, nil
}

// 实现 AssetLibrarySerivceServer 接口中的 UpdateAsset 方法
func (s *AssetLibraryService) UpdateAsset(ctx context.Context, req *services.UpdateAssetRequest) (*services.UpdateAssetResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	// TODO: 实现具体的修改逻辑
	entry := &entries.UpdateAssetEntry{
		Title:         req.Title,
		Description:   req.Description,
		Tags:          req.Tags,
		Thumbnail:     req.Thumbnail,
		ThumbnailSize: req.ThumbnailSize,
	}
	if external, err := convertToMap(req.External); err == nil {
		entry.External = external
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	if err := logic.UpdateAsset(req.AppId, req.UserId, req.Id, req.PermissionType, entry); err != nil {
		errCode := SendErrorToJson(err)
		return &services.UpdateAssetResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}

	// 返回一个 UpdateAssetResponse 实例和 nil 错误
	return &services.UpdateAssetResponse{
		Code:    200,
		Message: "success",
	}, nil
}

// 实现 AssetLibrarySerivceServer 接口中的 DeleteAsset 方法
func (s *AssetLibraryService) DeleteAsset(ctx context.Context, req *services.DeleteAssetRequest) (*services.DeleteAssetResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	// TODO: 实现具体的删除逻辑
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.DeleteAssetResponse{
			Code: 50105,
		}, nil
	}
	if err := logic.DeleteAsset(*commonRequest.UserId, *commonRequest.AppId, req.Id, req.PermissionType); err != nil {
		errCode := SendErrorToJson(err)
		return &services.DeleteAssetResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 DeleteAssetResponse 实例和 nil 错误
	return &services.DeleteAssetResponse{
		Code:    200,
		Message: "success",
	}, nil
}

// 实现 AssetLibrarySerivceServer 接口中的 GetAsset 方法
func (s *AssetLibraryService) GetAsset(ctx context.Context, req *services.GetAssetRequest) (*services.GetAssetResponse, error) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		xutil.Logger.Errorf("panic: %v", r)
	//	}
	//}()
	// TODO: 实现具体的查询逻辑
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.GetAssetResponse{
			Code: 50105,
		}, nil
	}
	assetDto, err := logic.GetAsset(*commonRequest.UserId, *commonRequest.AppId, req.Id, req.PermissionType)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.GetAssetResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	asset := &services.Asset{
		Id:                   assetDto.Id,
		FileName:             assetDto.FileName,
		FileSize:             assetDto.FileSize,
		Type:                 assetDto.Type,
		Class:                assetDto.Class,
		ClassName:            assetDto.ClassName,
		Tags:                 assetDto.Tags,
		TagsName:             assetDto.TagsName,
		Title:                assetDto.Title,
		InstanceId:           assetDto.InstanceId,
		AppId:                assetDto.AppId,
		Description:          assetDto.Description,
		Thumbnail:            assetDto.Thumbnail,
		ThumbnailDownloadUrl: assetDto.ThumbnailDownloadUrl,
		Url:                  assetDto.Url,
		AssetDownloadUrl:     assetDto.AssetDownloadUrl,
		PermissionType:       assetDto.PermissionType,
		Status:               int32(assetDto.Status),
		Source:               assetDto.Source,
		Creator:              assetDto.Creator,
		CreatorName:          assetDto.CreatorName,
		Owner:                assetDto.Owner,
		RelatedAssets:        assetDto.RelatedAssets,
		CreateTime:           assetDto.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:           assetDto.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if assetDto.External != nil && len(assetDto.External) > 0 {
		external, err := convertToStructMap(assetDto.External)
		if err != nil {
			asset.External = external
		}
	}
	res := &services.GetAssetResponse{
		Code:    200,
		Message: "success",
		Asset:   asset,
	}
	// 返回一个 GetAssetResponse 实例和 nil 错误
	return res, nil
}

// 实现 AssetLibrarySerivceServer 接口中的 GetAssetList 方法
func (s *AssetLibraryService) UpdateAssetThumbnail(ctx context.Context, req *services.UpdateAssetThumbnailRequest) (*services.UpdateAssetThumbnailResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic: %v", r)
		}
	}()
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.UpdateAssetThumbnailResponse{
			Code: 50105,
		}, nil
	}
	entry := &entries.UpdateAssetThumbnailEntry{
		ThumbnailSize: req.ThumbnailSize,
	}
	// validate request
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	thumbnail, err := logic.UpdateAssetThumbnail(*commonRequest.AppId, *commonRequest.UserId, req.AssetId, req.PermissionType, entry)
	if err != nil {
		errCode := SendErrorToJson(err)

		return &services.UpdateAssetThumbnailResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 UpdateAssetThumbnailResponse 实例和 nil 错误
	return &services.UpdateAssetThumbnailResponse{
		Code:      200,
		Message:   "success",
		Thumbnail: thumbnail,
	}, nil
}
func (s *AssetLibraryService) AddClass(ctx context.Context, req *services.AddClassRequest) (*services.AddClassResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.AppId == nil {
		return &services.AddClassResponse{
			Code: 50105,
		}, nil
	}
	entry := &entries.CreateClassEntry{
		Name:        req.Name,
		Description: req.Description,
		Source:      req.Source,
		Creator:     *commonRequest.UserId,
		Types:       req.Types,
		MaxSize:     req.MaxSize,
		//PeimitInstanceID: req.PeimitInstanceID,
		IsTranfrom: req.IsTranfrom,
		Pipelines:  req.Pipelines,
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	classId, err := logic.AddClass(ctx, *commonRequest.AppId, entry)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.AddClassResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 AddClassResponse 实例和 nil 错误
	return &services.AddClassResponse{
		Code:    200,
		Message: "success",
		ClassId: classId,
	}, nil
}
func (s *AssetLibraryService) UpdateClass(ctx context.Context, req *services.UpdateClassRequest) (*services.UpdateClassResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.UpdateClassResponse{
			Code: 50105,
		}, nil
	}
	entry := &entries.UpdateClassEntry{
		Name:        req.Name,
		Description: req.Description,
		Types:       req.Types,
		MaxSize:     req.MaxSize,
		//PeimitInstanceID: req.PeimitInstanceID,
		IsTranfrom: req.IsTranfrom,
		Pipelines:  req.Pipelines,
	}
	err := logic.UpdateClass(ctx, req.ClassId, entry)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.UpdateClassResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 UpdateClassResponse 实例和 nil 错误
	return &services.UpdateClassResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) DeleteClass(ctx context.Context, req *services.DeleteClassRequest) (*services.DeleteClassResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.DeleteClassResponse{
			Code: 50105,
		}, nil
	}
	err := logic.DeleteClass(ctx, req.ClassId)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.DeleteClassResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 DeleteClassResponse 实例和 nil 错误
	return &services.DeleteClassResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) QueryClass(ctx context.Context, req *services.QueryClassRequest) (*services.QueryClassResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.AppId == nil {
		return &services.QueryClassResponse{
			Code: 50105,
		}, nil
	}
	classes, err := logic.QueryClass(ctx, *commonRequest.AppId)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.QueryClassResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 QueryClassResponse 实例和 nil 错误
	QueryClassResponse := &services.QueryClassResponse{
		Code:    200,
		Message: "success",
	}
	ClassesProto := make([]*services.Class, len(classes))
	for i, class := range classes {
		ClassesProto[i] = &services.Class{
			ClassId:     class.ClassId,
			Name:        class.Name,
			Description: class.Description,
			Source:      class.Source,
			Creator:     class.Creator,
			Types:       class.Types,
			MaxSize:     class.MaxSize,
			//PeimitInstanceID: class.PeimitInstanceID,
			IsTranfrom: class.IsTranfrom,
			//Pipelines:  class.Pipelines,
		}
	}
	QueryClassResponse.Classes = ClassesProto
	return QueryClassResponse, nil
}
func (s *AssetLibraryService) GetClass(ctx context.Context, req *services.GetClassRequest) (*services.GetClassResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.AppId == nil {
		return &services.GetClassResponse{
			Code: 50105,
		}, nil
	}
	class, err := logic.GetClass(ctx, req.ClassId)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.GetClassResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 GetClassResponse 实例和 nil 错误
	return &services.GetClassResponse{
		Code:    200,
		Message: "success",
		Class: &services.Class{
			ClassId:     class.ClassId,
			Name:        class.Name,
			Description: class.Description,
			Source:      class.Source,
			Creator:     class.Creator,
			Types:       class.Types,
			MaxSize:     class.MaxSize,
			//PeimitInstanceID: class.PeimitInstanceID,
			IsTranfrom: class.IsTranfrom,
			//Pipelines:  class.Pipelines,
		},
	}, nil
}
func (s *AssetLibraryService) AddTag(ctx context.Context, req *services.AddTagRequest) (*services.AddTagResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.AddTagResponse{
			Code: 50105,
		}, nil
	}
	entry := &entries.CreateTagEntry{
		Name:        req.Name,
		Description: req.Description,
		AppId:       req.AppId,
		ClassId:     req.ClassId,
		Source:      req.Source,
		Creator:     *commonRequest.UserId,
		Owner:       req.Owner,
		Type:        req.Type,
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	tagId, err := logic.AddTag(ctx, entry)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.AddTagResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 AddTagResponse 实例和 nil 错误
	return &services.AddTagResponse{
		Code:    200,
		Message: "success",
		TagId:   tagId}, nil
}
func (s *AssetLibraryService) UpdateTag(ctx context.Context, req *services.UpdateTagRequest) (*services.UpdateTagResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.UserId == nil {
		return &services.UpdateTagResponse{
			Code: 50105,
		}, nil
	}
	entry := &entries.UpdateTagEntry{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := logic.UpdateTag(ctx, req.Id, entry); err != nil {
		errCode := SendErrorToJson(err)
		return &services.UpdateTagResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 UpdateTagResponse 实例和 nil 错误
	return &services.UpdateTagResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) DeleteTag(ctx context.Context, req *services.DeleteTagRequest) (*services.DeleteTagResponse, error) {
	if err := logic.DeleteTag(ctx, req.Id); err != nil {
		errCode := SendErrorToJson(err)
		return &services.DeleteTagResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 DeleteTagResponse 实例和 nil 错误
	return &services.DeleteTagResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) QueryTag(ctx context.Context, req *services.QueryTagRequest) (*services.QueryTagResponse, error) {
	commonRequest := req.CommonRequest
	//if commonRequest == nil || commonRequest.AppId == nil {
	//	return &services.QueryTagResponse{
	//		Code: 50105,
	//	}, nil
	//}
	tags, err := logic.QueryTag(ctx, commonRequest.AppId, req.ClassId, req.Name, req.Owner, req.Type)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.QueryTagResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	tagsDto := make([]*services.Tag, len(tags))
	for i, tag := range tags {
		tagsDto[i] = &services.Tag{
			Id:          tag.TagId,
			Name:        tag.Name,
			Description: tag.Description,
			AppId:       tag.AppId,
			ClassId:     tag.ClassId,
			Source:      tag.Source,
			UserId:      tag.Creator,
			Owner:       tag.Owner,
			Type:        tag.Type,
			CreateTime:  tag.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  tag.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}
	// 返回一个 QueryTagResponse 实例和 nil 错误
	return &services.QueryTagResponse{
		Code:    200,
		Message: "success",
		Tag:     tagsDto,
	}, nil
}
func (s *AssetLibraryService) QueryTagList(ctx context.Context, req *services.QueryTagListRequest) (*services.QueryTagListResponse, error) {
	commonRequest := req.CommonRequest
	if commonRequest == nil || commonRequest.AppId == nil {
		return &services.QueryTagListResponse{
			Code: 50105,
		}, nil
	}
	queryTagReq := &httpDto.QueryTagReq{
		Owner: req.Owner,
		Type:  req.Type,
	}
	tags, err := logic.QueryTagList(ctx, *commonRequest.AppId, queryTagReq)
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.QueryTagListResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	tagsDto := make([]*services.Tag, len(tags))
	for i, tag := range tags {
		tagsDto[i] = &services.Tag{
			Id:          tag.TagId,
			Name:        tag.Name,
			Description: tag.Description,
			AppId:       tag.AppId,
			ClassId:     tag.ClassId,
			Source:      tag.Source,
			UserId:      tag.Creator,
			Owner:       tag.Owner,
			Type:        tag.Type,
			CreateTime:  tag.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  tag.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}
	// 返回一个 QueryTagListResponse 实例和 nil 错误
	return &services.QueryTagListResponse{
		Code:    200,
		Message: "success",
		Tag:     tagsDto,
	}, nil
}
func (s *AssetLibraryService) AddInstanceScopeId(ctx context.Context, req *services.AddInstanceScopeIdRequest) (*services.AddInstanceScopeIdResponse, error) {
	entry := &entries.CreateInstanceScopeIdEntry{
		InstanceId: req.InstanceId,
		ScopeIdMap: req.ScopeIdMap,
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	if _, err := logic.AddInstanceScopeId(entry); err != nil {
		errCode := SendErrorToJson(err)
		return &services.AddInstanceScopeIdResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 AddInstanceScopeIdResponse 实例和 nil 错误
	return &services.AddInstanceScopeIdResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) DeleteInstanceScopeId(ctx context.Context, req *services.DeleteInstanceScopeIdRequest) (*services.DeleteInstanceScopeIdResponse, error) {
	if err := logic.DeleteInstanceScopeId(req.InstanceId); err != nil {
		errCode := SendErrorToJson(err)
		return &services.DeleteInstanceScopeIdResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 DeleteInstanceScopeIdResponse 实例和 nil 错误
	return &services.DeleteInstanceScopeIdResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) GetInstanceScopeId(ctx context.Context, req *services.GetInstanceScopeIdRequest) (*services.GetInstanceScopeIdResponse, error) {
	instanceScopeId_DB, err := logic.GetInstanceScopeId(req.InstanceId)
	errCode := SendErrorToJson(err)
	if err != nil {
		return &services.GetInstanceScopeIdResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 GetInstanceScopeIdResponse 实例和 nil 错误
	return &services.GetInstanceScopeIdResponse{
		Code:    200,
		Message: "success",
		InstanceScopeId: &services.InstanceScopeId{
			InstanceId: instanceScopeId_DB.InstanceId,
			ScopeIdMap: instanceScopeId_DB.ScopeIdMap,
		},
	}, nil
}
func (s *AssetLibraryService) UpdateInstanceScopeId(ctx context.Context, req *services.UpdateInstanceScopeIdRequest) (*services.UpdateInstanceScopeIdResponse, error) {
	entry := &entries.UpdateInstanceScopeIdEntry{
		InstanceId: req.InstanceId,
		ScopeIdMap: req.ScopeIdMap,
	}
	if err := router.ValidateJSON(entry); err != nil {
		err = router.CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	if err := logic.UpdateInstanceScopeId(entry); err != nil {
		errCode := SendErrorToJson(err)
		return &services.UpdateInstanceScopeIdResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	// 返回一个 UpdateInstanceScopeIdResponse 实例和 nil 错误
	return &services.UpdateInstanceScopeIdResponse{
		Code:    200,
		Message: "success",
	}, nil
}
func (s *AssetLibraryService) QueryInstanceScopeId(ctx context.Context, req *services.QueryInstanceScopeIdRequest) (*services.QueryInstanceScopeIdResponse, error) {
	instanceScopeId_DBs, err := logic.QueryInstanceScopeIdList()
	if err != nil {
		errCode := SendErrorToJson(err)
		return &services.QueryInstanceScopeIdResponse{
			Code:    errCode,
			Message: err.Error(),
		}, err
	}
	InstanceScopeIds := make([]*services.InstanceScopeId, len(instanceScopeId_DBs))
	for i, instanceScopeId_DB := range instanceScopeId_DBs {
		InstanceScopeIds[i] = &services.InstanceScopeId{
			InstanceId: instanceScopeId_DB.InstanceId,
			ScopeIdMap: instanceScopeId_DB.ScopeIdMap,
		}
	}
	// 返回一个 QueryInstanceScopeIdResponse 实例和 nil 错误
	return &services.QueryInstanceScopeIdResponse{
		Code:             200,
		Message:          "success",
		InstanceScopeIds: InstanceScopeIds,
	}, nil
}
func SendErrorToJson(err error) int32 {
	errStatus, _ := status.FromError(err)

	var errDetails *services.ApiExceptionRpc
	// convert interface to map
	if errStatus.Details() != nil && len(errStatus.Details()) > 0 {
		errDetails = errStatus.Details()[0].(*services.ApiExceptionRpc)
	}
	fmt.Println("Err Details", int(errDetails.GetErrorDetail().GetCode()))
	errorCode := int(errDetails.GetErrorDetail().GetCode())
	// returning error to client using framework
	var errCode int32 = int32(errorCode)
	if errorCode == 0 {
		errorCode = 1001
	}
	return errCode
}
