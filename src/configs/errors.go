package configs

import (
	"gitlab.zixel.cn/go/framework"
)

var ErrorCodes = map[int]string{
	200:   "Success.",
	1000:  "Server internal error.",
	1001:  "Database error.",
	10005: "invalid primitive id ",
	10006: "document is nil",
	10007: "primitive id conversion error",
	10008: "Decoding Error",
	10009: "Failed to validate request body.",
	10010: "Invalid Config Type.",
	10011: "Data already exist.",
	10012: "Config not found.",
	10013: "No result found",
	10014: "Max size cannot be nil",
	10015: "Class not found.",
	10016: "Invalid tag Type",
	10017: "Invalid classId",
	10018: "Tag Not found",

	40001: "Failed to validate param.",

	50001: "Invalid body format.",
	50002: "Validate data failed.",
	50003: "Type is failed.",
	50100: "permissionType is invalid",
	50101: "tenantID is empty",
	50102: "checkPermission is failed",
	50103: "GetUploadFileUrl is error",
	50104: "GRPC is error",
	50105: "commonRequest is error",
	50106: "assetUploaded is error",
	50107: "QueryConfigByType is nil",
	50108: "File size exceeds asset library configuration",
	50109: "Number of files exceeds asset library configuration",
	50110: "The selected resource class does not support this file format!",
	50111: "File size exceeds class limit!",
	50112: "MIME Type is failed",
	50113: "FileHashName error",
	50114: "Failed to get user by openId",
	50115: "User does not exist",
	50116: "thumbnail size or type empty",
	50117: "Invalid assetId",
	50118: "No Asset Found",
	50119: "Please do not make repeated requests",
	50120: "Post convert job error",
	50121: "get current working directory error",
	50122: "Get DownLoadSign Url error",
	50123: "Download file error",
	50124: "VideoGenSnapshot is error",
	50125: "Open file error",
	50126: "File.Stat error",
	50127: "Upload file error",
	50128: "Instance scope id not found",
	50129: "CopyObject error",
	50130: "ListUserByUid  is error",
}

type ErrorNo struct {
	Code    int
	Message string
	Err     error
}

func (err *ErrorNo) ErrCode() int {
	if err.Code == 0 {
		return 1001
	}
	return err.Code
}

func (err *ErrorNo) Error() error {
	if err.Err == nil {
		if err.Message == "" {
			err.Message = ErrorCodes[err.Code]
		}
		err.Err = framework.NewServiceError(err.Code, err.Message)
	}
	return err.Err
}
