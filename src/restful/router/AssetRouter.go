package router

import (
	"assetLibary/configs"
	"assetLibary/restful/contants"
	"assetLibary/restful/entries"
	"assetLibary/restful/handler"
	"assetLibary/restful/httpDto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func AddAsset(c *gin.Context) {
	var (
		err    error
		result *entries.AddAssetRpn
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userID := c.GetHeader(contants.ZixelHeader_UserId)

	entry := &entries.CreateAssetEntry{}
	if err = c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = ValidateJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	entry.User = userID
	result, err = handler.AddAsset(ctx, appId, entry)
}
func UpdateAsset(c *gin.Context) {
	var (
		err    error
		result *entries.AddAssetRpn
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userID := c.GetHeader(contants.ZixelHeader_UserId)
	paramAsset := &entries.ParamAssetUri_t{}
	if err = c.BindUri(paramAsset); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	entry := &entries.UpdateAssetEntry{}
	if err := c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = ValidateJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err = handler.UpdateAsset(ctx, appId, userID, paramAsset.AssetId, paramAsset.PermissionType, entry)
}

func UpdateAssetThumbnail(c *gin.Context) {
	var (
		err    error
		result *httpDto.UpdateThumbnailRpn
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userID := c.GetHeader(contants.ZixelHeader_UserId)
	paramAsset := &entries.ParamAssetUri_t{}
	if err = c.BindUri(paramAsset); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	entry := &entries.UpdateAssetThumbnailEntry{}
	if err = ValidateJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err := c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// validate request
	result, err = handler.UpdateAssetThumbnail(ctx, appId, userID, paramAsset.AssetId, paramAsset.PermissionType, entry)
}

func AssetUploaded(c *gin.Context) {
	var (
		err    error
		result *entries.AddAssetRpn
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	paramAsset := &entries.ParamAssetUri_t{}
	if err = c.BindUri(paramAsset); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = handler.AssetUploaded(ctx, appId, paramAsset.AssetId, paramAsset.PermissionType)
}

func QueryAsset(c *gin.Context) {
	var (
		err    error
		result *httpDto.GetAssetListRpn_t
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	permissionType := c.Param("permissionType")
	entry := &entries.GetAssetListReq_t{}
	if err := c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	entry.PermissionType = permissionType

	result, err = handler.QueryAsset(ctx, appId, userId, permissionType, entry)
}

func DeleteAsset(c *gin.Context) {
	var (
		err    error
		result *httpDto.GetAssetListRpn_t
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	entry := &entries.ParamAssetUri_t{}
	if err = c.BindUri(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err = handler.DelateAsset(ctx, userId, appId, entry.AssetId, entry.PermissionType)
}

func GetAsset(c *gin.Context) {
	var (
		err    error
		result *httpDto.AssetDto
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	entry := &entries.ParamAssetUri_t{}
	if err = c.BindUri(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	result, err = handler.GetAsset(ctx, userId, appId, entry.AssetId, entry.PermissionType)
}

func UseAsset(c *gin.Context) {
	var (
		err    error
		result string
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	entry := &entries.ParamAssetUri_t{}
	if err = c.BindUri(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	copyReq := &entries.CopyAssetReq_t{}
	if err := c.ShouldBindJSON(copyReq); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = ValidateJSON(copyReq); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	result, err = handler.UseAsset(ctx, userId, appId, entry.AssetId, entry.PermissionType, copyReq)

}

func CheckPermission(c *gin.Context) {
	var (
		err    error
		result map[string]bool
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	checkPermissionReq_t := &entries.CheckPermissionReq_t{}
	if err = c.ShouldBindJSON(checkPermissionReq_t); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if err = ValidateJSON(checkPermissionReq_t); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	result, err = handler.CheckPermission(ctx, userId, appId, checkPermissionReq_t)
}
