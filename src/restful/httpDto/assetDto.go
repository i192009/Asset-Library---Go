package httpDto

import "time"

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
