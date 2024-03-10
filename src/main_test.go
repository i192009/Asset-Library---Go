package main

import (
	"assetLibary/restful/entries"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

type AddAssetRpn struct {
	AssetId        string            `json:"assetId" binding:"required"`
	AssetUploadUrl string            `json:"assetUploadUrl" binding:"required"`
	ThumbUploadUrl string            `json:"thumbUploadUrl,omitempty"`
	OtherFile      map[string]string `json:"otherFile,omitempty"`
}

type GetAssetListRpn_t struct {
	Page    int64       `json:"page" binding:"required"`
	Size    int64       `json:"pagesize" binding:"required"`
	Total   int64       `json:"totalCount" binding:"required"`
	Results []*AssetDto `json:"results"`
}
type AssetDto struct {
	Id                   string                 `bson:"_id" json:"assetId"`
	FileName             string                 `bson:"filename" json:"fileName"`
	FileSize             int64                  `bson:"filesize" json:"fileSize"`
	Type                 string                 `bson:"type"  json:"type"`
	Class                string                 `bson:"class,omitempty" json:"class"`
	ClassName            string                 `bson:"className,omitempty" json:"className"`
	Tags                 []string               `bson:"tags,omitempty" json:"tags"`
	TagsName             []string               `bson:"tagsNames,omitempty" json:"tagsNames"`
	Title                string                 `bson:"title" json:"title"`
	InstanceId           string                 `bson:"instanceId" json:"instanceId"`
	AppId                string                 `bson:"appId" json:"appId"`
	Description          string                 `bson:"description,omitempty" json:"description"`
	Thumbnail            string                 `bson:"thumbnail,omitempty" json:"thumbnail"`
	ThumbnailDownloadUrl string                 `bson:"thumbnailDownloadUrl,omitempty" json:"thumbnailDownloadUrl"`
	Url                  string                 `bson:"url" json:"url"`
	AssetDownloadUrl     string                 `bson:"assetDownloadUrl" json:"assetDownloadUrl"`
	PermissionType       string                 `bson:"permissionType" json:"permissionType"` //资产权限类型。1，共有  2，企业 3，私有
	Status               int                    `bson:"status" json:"status"`                 //资产状态 New Normal Reject Deleted
	Source               string                 `bson:"source" json:"source"`
	External             map[string]interface{} `bson:"external,omitempty" json:"external"` /// 额外的数据，供开发者保存一些自定义的资产属性
	Creator              string                 `bson:"creator" json:"creator"`
	CreatorName          string                 `bson:"creatorName" json:"creatorName"`
	Owner                string                 `bson:"owner" json:"owner"`
	RelatedAssets        []string               `bson:"relatedAssets,omitempty" json:"relatedAssets"` //关联资产
	CreateTime           time.Time              `bson:"createTime" json:"createTime"`
	UpdateTime           time.Time              `bson:"updateTime" json:"updateTime"`
}

type UpdateThumbnailRpn struct {
	Thumbnail string `json:"thumbnail"`
}

func Add(fileName string, size int64) *AddAssetRpn {
	fmt.Println("AddAsset" + fileName)
	url := "https://dev.zixel.cn/api/assetManage/v1/backend/asset/add"
	method := "POST"
	ent := entries.CreateAssetEntry{

		FileName:       fileName,
		FileSize:       size,
		Type:           "application/fbx",
		Class:          "647ea975b3fc4fd31de4a691",
		Title:          fileName,
		PermissionType: "1",
		Owner:          "un_941318144",
		Description:    "test zblock user lib",
		Tags:           []string{"647ea9a7b3fc4fd31de4a692"},
	}
	by, _ := json.Marshal(ent)
	payload := bytes.NewReader(by)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Zixel-Application-Id", "1002")
	req.Header.Add("Zixel-Open-Id", "oi_aba23d0e8e484830b6ac41c83fe4faf0")
	req.Header.Add("Zixel-Auth-Token", "ut_1TqEd85RMjhN5SK9O7pW31yOUm0aX4aQ")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.zixel.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rpn := &AddAssetRpn{}
	json.Unmarshal(body, rpn)
	fmt.Println(string(body))
	return rpn
}

func Upload(assetId string) {
	url := "https://dev.zixel.cn/api/assetManage/v1/backend/asset/assetUploaded/1/" + assetId
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Zixel-Application-Id", "1002")
	req.Header.Add("Zixel-Open-Id", "oi_aba23d0e8e484830b6ac41c83fe4faf0")
	req.Header.Add("Zixel-Auth-Token", "ut_1TqEd85RMjhN5SK9O7pW31yOUm0aX4aQ")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.zixel.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Upload+++++++:" + string(body))
}
func TestUploadTh(t *testing.T) {
	dir := "C:\\Users\\zhanghao\\Desktop\\BoltA35\\tools" // 要读取的目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("读取目录失败", err)
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() { // 如果不是目录
			data, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				fmt.Println("读取文件失败", err)
				os.Exit(1)
			}
			fileName := filepath.Base(file.Name())
			fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

			GetAssetListRpn_t := GetAsset(fileNameWithoutExt + ".fbx")
			if GetAssetListRpn_t != nil && len(GetAssetListRpn_t.Results) > 0 {
				UpdateThumbnailRpn := Thumbnail(GetAssetListRpn_t.Results[0].Id, file.Size())
				UploadTh(UpdateThumbnailRpn.Thumbnail, string(data))
			}
		}
	}
}

func UploadTh(url string, data string) {
	fmt.Println("UploadTh+++++++++++++++++++++++++++")
	method := "PUT"
	//payload := strings.NewReader("<file contents here>")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(data))

	if err != nil {
		fmt.Println("------4")
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Host", "re-link.obs.cn-east-3.myhuaweicloud.com")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Type", "application/octet-stream")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("------3")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("------2")

		fmt.Println(err)
		return
	}
	fmt.Println("------1")

	fmt.Println(string(body))
}
func Thumbnail(assetId string, size int64) *UpdateThumbnailRpn {
	url := "https://qa.zixel.cn/api/assetManage/v1/backend/asset/update/1/" + assetId + "/thumbnail"
	method := "PUT"

	payload := strings.NewReader(`{` + `
	"thumbnailSize": ` + strconv.FormatInt(size, 10) + `
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Zixel-Application-Id", "1002")
	req.Header.Add("Zixel-Open-Id", "oi_12919898112")
	req.Header.Add("Zixel-Auth-Token", "ut_4wm3smdnh6t2TGY4gXCv94ywXM6PDKrV")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "qa.zixel.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	UpdateThumbnailRpn := &UpdateThumbnailRpn{}
	json.Unmarshal(body, UpdateThumbnailRpn)
	fmt.Println(string(body))
	return UpdateThumbnailRpn
}

func GetAsset(fileName string) *GetAssetListRpn_t {
	fmt.Println("GetAsset+++++++++++++++++++++++++++" + fileName)
	url := "https://qa.zixel.cn/api/assetManage/v1/backend/asset/1/query"
	method := "POST"

	payload := strings.NewReader(`{` + `
	"page": 0,` + `
	"pagesize": 10,` + `
	"search":"` + fileName + `",` + `
	"sort":-1` + `
}`)
	fmt.Println("++++++++++++++++++++++++")

	fmt.Println(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Zixel-Application-Id", "1002")
	req.Header.Add("Zixel-Open-Id", "oi_12919898112")
	req.Header.Add("Zixel-Auth-Token", "ut_4wm3smdnh6t2TGY4gXCv94ywXM6PDKrV")
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "qa.zixel.cn")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rpn := &GetAssetListRpn_t{}
	json.Unmarshal(body, rpn)
	fmt.Println(string(body))
	return rpn
}
func TestUpload(t *testing.T) {

	dir := "C:\\Users\\zhanghao\\Desktop\\tools" // 要读取的目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("读取目录失败", err)
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() { // 如果不是目录
			data, err := ioutil.ReadFile(dir + "/" + file.Name())
			if err != nil {
				fmt.Println("读取文件失败", err)
				os.Exit(1)
			}

			addRpn := Add(file.Name(), file.Size())
			if addRpn == nil {
				fmt.Println("上传失败")
				continue
			}
			method := "PUT"
			url := addRpn.AssetUploadUrl
			client := &http.Client{}
			req, err := http.NewRequest(method, url, strings.NewReader(string(data)))
			fmt.Println(string(data)) // 打印文件内容
			if err != nil {
				fmt.Println("------4")
				fmt.Println(err)
				return
			}
			req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
			req.Header.Add("Accept", "*/*")
			req.Header.Add("Sec-Fetch-Mode", "cors")
			req.Header.Add("Sec-Fetch-Dest", "empty")
			req.Header.Add("Sec-Fetch-Site", "cross-site")
			req.Header.Add("Host", "re-link.obs.cn-east-3.myhuaweicloud.com")
			req.Header.Add("Connection", "keep-alive")
			//req.Header.Add("Content-Type", "application/octet-stream")

			res, err := client.Do(req)
			if err != nil {
				fmt.Println("------3")
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("------2")

				fmt.Println(err)
				return
			}
			fmt.Println("------1")

			fmt.Println(string(body))
			Upload(string(addRpn.AssetId))

		}
		//file, err := os.Open("C:\\Users\\zhanghao\\Desktop\\tools\\钢丝钳.fbx")
		//if err != nil {
		//	panic(err)
		//}
		//defer file.Close()
		//payload := strings.NewReader("<file contents here>")

	}

	//url := "https://zetaverse.obs.cn-east-3.myhuaweicloud.com:443/qa/1002/2f/d9/53/bfb4/8fce306fd66eee996d7a7c.fbx?AccessKeyId=VKQA5OEEYYYF3XE7JSAD&Expires=1686105143&Signature=oUSldVJ07kZf1aZ%2Bv48jLyF05UM%3D"
	//method := "PUT"
	//
	//file, err := os.Open("C:\\Users\\zhanghao\\Desktop\\tools\\钢丝钳.fbx")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	////payload := strings.NewReader("<file contents here>")
	//
	//client := &http.Client{
	//}
	//req, err := http.NewRequest(method, url, file)
	//
	//if err != nil {
	//	fmt.Println("------4")
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Sec-Fetch-Mode", "cors")
	//req.Header.Add("Sec-Fetch-Dest", "empty")
	//req.Header.Add("Sec-Fetch-Site", "cross-site")
	//req.Header.Add("Host", "re-link.obs.cn-east-3.myhuaweicloud.com")
	//req.Header.Add("Connection", "keep-alive")
	////req.Header.Add("Content-Type", "application/octet-stream")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("------3")
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("------2")
	//
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("------1")
	//
	//fmt.Println(string(body))

}
func TestStrings(t *testing.T) {
	//if strings.Contains("application/fbx", contants.FileType_File) {
	//	fmt.Println("true")
	//}
	fmt.Println(strings.TrimSuffix("dev/uuu.png", filepath.Ext("dev/uuu.png")))
}

func TestUp(t *testing.T) {

	//url := "https://zetaverse.obs.cn-east-3.myhuaweicloud.com:443/qa/1002/5d/8b/ce/3153/9ca45db9f1f7da54e6688f.wav2?AccessKeyId=VKQA5OEEYYYF3XE7JSAD&Expires=1688043225&Signature=oJMRdxEBeZLytk82j447weHazSo%3D"
	url := "https://zetaverse.obs.cn-east-3.myhuaweicloud.com:443/dev/e6/2d/57/8a99/a9c8ce85260be39c6ab049.thumb?AccessKeyId=VKQA5OEEYYYF3XE7JSAD&Expires=1693205173&Signature=meNK61UjbyE5f%2BUvlWuhb5OG6RE%3D"
	method := "PUT"

	file, err := os.Open("C:\\Users\\zhanghao\\Desktop\\BoltA35\\1685603543570.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//payload := strings.NewReader("<file contents here>")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, file)

	if err != nil {
		fmt.Println("------4")
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	//req.Header.Add("Host", "re-link.obs.cn-east-3.myhuaweicloud.com")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Type", "audio/wav")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("------3")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("------2")

		fmt.Println(err)
		return
	}
	fmt.Println("------1")

	fmt.Println(string(body))
}
