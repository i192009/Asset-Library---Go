package contants

var (
	ContentType_AssetType             = []string{"image", "image", "audio", "video", "text", "application", "application"}
	ContentType_AssetSubType_Audio    = []string{"mp4", "mp3", "mp4", "ogg"}
	ContentType_AssetSubType_Video    = []string{"mp4", "x-msvideo", "mp4"}
	ContentType_AssetSubType_Pic      = []string{"png", "jpeg", "png", "svg", "gif"}
	ContentType_AssetSubType_Text     = []string{"plain", "plain", "yaml", "toml", "html"}
	ContentType_AssetSubType_RichText = []string{"json", "rtf", "json", "msword", "vnd.ms-excel", "xml"}
	ContentType_AssetSubType_File     = []string{"bin", "zip", "x-tar", "x-bzip-compressed-tar", "x-compressed-tar", "octet-stream", "octet-stream"}
)

var ContentType = map[int][]string{
	AssetType_Pic:      ContentType_AssetSubType_Pic,
	AssetType_Audio:    ContentType_AssetSubType_Audio,
	AssetType_Video:    ContentType_AssetSubType_Video,
	AssetType_Text:     ContentType_AssetSubType_Text,
	AssetType_RichText: ContentType_AssetSubType_RichText,
	AssetType_File:     ContentType_AssetSubType_File,
}
var (
	AssetStatus_New     = 1
	AssetStatus_Review  = 2
	AssetStatus_Normal  = 3
	AssetStatus_Reject  = 4
	AssetStatus_Deleted = 5
)

var (
	AssetType_Pic       = 1
	AssetType_Audio     = 2
	AssetType_Video     = 3
	AssetType_Text      = 4
	AssetType_RichText  = 5
	AssetType_File      = 6
	AssetType_ABPackage = 7
)

// ContentType_AssetType             = []string{"image", "image", "audio", "video", "text", "application", "application"}
var (
	FileType_Pic   = "image"
	FileType_Audio = "audio"
	FileType_Video = "video"
	FileType_Text  = "text"
	FileType_File  = "application"
)

var (
	AssetSubType_Pic_JPG = 1
	AssetSubType_Pic_PNG = 2
	AssetSubType_Pic_SVG = 3
	AssetSubType_Pic_GIF = 4
)

var (
	AssetSubType_Audio_MP3 = 1
	AssetSubType_Audio_MP4 = 2
	AssetSubType_Audio_OGG = 3
)

var (
	AssetSubType_Video_AVI = 1
	AssetSubType_Video_MP4 = 2
)

var (
	AssetSubType_Text_Plain = 1
	AssetSubType_Text_YAML  = 2
	AssetSubType_Text_TOML  = 3
	AssetSubType_Text_HTML  = 4
)

var (
	AssetSubType_RichText_RTF  = 1
	AssetSubType_RichText_JSON = 2
	AssetSubType_RichText_DOC  = 3
	AssetSubType_RichText_XLS  = 4
	AssetSubType_RichText_XML  = 5
)

var (
	AssetSubType_File_ZIP = 1
	AssetSubType_File_TAR = 2
	AssetSubType_File_BZ  = 3
	AssetSubType_File_GZ  = 4
	AssetSubType_File_BIN = 5
	AssetSubType_File_3D  = 6
)

var (
	AssetPermissionType_Public  = "1"
	AssetPermissionType_Tenant  = "2"
	AssetPermissionType_Private = "3"
)

var (
	OperateType_Read   = "1"
	OperateType_Write  = "2"
	OperateType_Delete = "3"
)

var (
	AssetPlatform_PC      = 1
	AssetPlatform_Android = 2
	AssetPlatform_IOS     = 3
)

var (
	ZixelHeader_UserId    string = "Zixel-User-Id"
	ZixelHeader_OpenId    string = "Zixel-Open-Id"
	ZixelHeader_AuthToken string = "Zixel-Auth-Token"
	ZixelHeader_AppId     string = "Zixel-Application-Id"
)
