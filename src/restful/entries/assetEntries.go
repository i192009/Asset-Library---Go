package entries

type CreateAssetEntry struct {
	FileName       string                 `json:"filename"`
	FileSize       int64                  `json:"filesize"`
	Type           string                 `json:"type"`
	Tags           []string               `json:"tags"`
	Class          string                 `json:"class"`
	Platform       string                 `json:"platform"`                           //平台。1，pc  2，android  3，ios
	Platforms      []string               `json:"platform"`                           //平台。1，pc  2，android  3，ios
	PermissionType string                 `json:"permissionType" validate:"required"` //资产权限类型。1，共有  2，企业 3，私有
	Title          string                 `json:"title" validate:"required,max=100"`
	AppId          string                 `json:"appId"`
	InstanceId     string                 `json:"instanceId"`
	Description    string                 `json:"description"  validate:"max=500"`
	External       map[string]interface{} `json:"external"` /// 额外的数据，供开发者保存一些自定义的资产属性
	Thumbnail      string                 `json:"thumbnail"`
	ThumbnailSize  int64                  `json:"thumbnailSize"`
	ThumbnailType  string                 `json:"thumbnailType"`
	OtherFile      map[string]AssetFile   `json:"otherFile"`
	User           string                 `json:"user"`
	Owner          string                 `json:"owner" validate:"required"`
	Source         string                 `json:"source"`
}
type AssetFile struct {
	Name    string `json:"name"`
	Type    int    `json:"type"`
	SubType int    `json:"subtype"`
	Size    int64  `json:"size"`
}

type AddAssetRpn struct {
	AssetId        string            `json:"assetId" binding:"required"`
	AssetUploadUrl string            `json:"assetUploadUrl" binding:"required"`
	ThumbUploadUrl string            `json:"thumbUploadUrl,omitempty"`
	OtherFile      map[string]string `json:"otherFile,omitempty"`
}
type GetAssetListReq_t struct {
	Page           int64  `form:"page" json:"page"`
	Size           int64  `form:"pagesize" json:"pagesize"`
	PageSize       int64  `form:"pageSize" json:"pageSize"`
	Type           int    `form:"type" json:"type"`
	Tags           string `form:"tags" json:"tags"`
	Platform       string `form:"platform" json:"platform"`             //平台。1，pc  2，android  3，ios
	PermissionType string `form:"permissionType" json:"permissionType"` //资产权限类型。1，共有  2，私有
	Class          string `form:"class" json:"class"`
	Order          string `form:"order" json:"order"`
	Sort           int    `form:"sort" json:"sort"`     //排序 0 - 升序， 1 - 降序
	Search         string `form:"search" json:"search"` //模糊查询字段
	Owner          string `form:"owner" json:"owner"`
}

type ParamAssetUri_t struct {
	PermissionType string `uri:"permissionType"`
	AssetId        string `uri:"assetId"`
}

type UpdateAssetEntry struct {
	Title         string                 `json:"title" validate:"max=100"`
	Description   string                 `json:"description" validate:"max=500"`
	External      map[string]interface{} `json:"external"` /// 额外的数据，供开发者保存一些自定义的资产属性
	Tags          []string               `json:"tags"`
	Thumbnail     string                 `json:"thumbnail"`
	ThumbnailSize int64                  `json:"thumbnailSize"`
}

type CopyAssetReq_t struct {
	TargetPosition string `json:"targetPosition" validate:"required"`
}

type CheckPermissionReq_t struct {
	PermissionType string `form:"permissionType" json:"permissionType" validate:"required"` //资产权限类型。1，共有 2，企业 3，私有
	AssetId        string `form:"assetId" json:"assetId"`
	TenantId       string `form:"tenantId" json:"tenantId"`
	OperateType    string `form:"operateType" json:"operateType"`
}

type UpdateAssetThumbnailEntry struct {
	ThumbnailSize int64  `json:"thumbnailSize" validate:"required"`
	ThumbnailType string `json:"thumbnailType"`
}
