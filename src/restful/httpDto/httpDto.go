package httpDto

import "assetLibary/restful/entries"

type QueryOperationRecordByIdReq struct {
	Id int `uri:"id" validate:"required"`
}
type SendMessageReq struct {
	Msg   string `uri:"msg" validate:"required"`
	Topic string `uri:"topic" validate:"required"`
}

type QueryOperationRecordReq struct {
	string
}

type AddAssetReq_t struct {
	FileName       string                       `json:"filename" validate:"required"`
	FileSize       int64                        `json:"filesize" validate:"required"`
	Type           int                          `json:"type" validate:"required"`
	SubType        int                          `json:"subtype" validate:"required"`
	Tags           []string                     `json:"tags"`
	AlbumId        string                       `json:"albumId"`
	Platform       string                       `json:"platform"`       //平台。1，pc  2，android  3，ios
	PermissionType string                       `json:"permissionType"` //资产权限类型。1，共有  2，私有
	Title          string                       `json:"title" validate:"required"`
	Description    string                       `json:"description"`
	External       map[string]interface{}       `json:"external"` /// 额外的数据，供开发者保存一些自定义的资产属性
	Thumbnail      string                       `json:"thumbnail"`
	ThumbnailSize  int64                        `json:"thumbnailSize"`
	ThumbnailType  int                          `json:"thumbnailType"`
	OtherFile      map[string]entries.AssetFile `json:"otherFile"`
}
type ParamAssetUri_t struct {
	AlbumId        string `uri:"albumId"`
	AssetId        string `uri:"assetId"`
	PermissionType string `uri:"permissionType"` //资产权限类型。1，共有  2，私有
}

type QueryAssetReq_t struct {
}
