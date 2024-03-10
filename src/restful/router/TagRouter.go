package router

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful/contants"
	"assetLibary/restful/entries"
	"assetLibary/restful/handler"
	"assetLibary/restful/httpDto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func AddTag(c *gin.Context) {
	var (
		err    error
		result *httpDto.CreateTagRpn
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
	userId := c.GetHeader(contants.ZixelHeader_UserId)
	appId := c.GetHeader(contants.ZixelHeader_AppId)
	entry := &entries.CreateTagEntry{}
	if err := c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		//log.Errorf(ctx, "AddTag is  error.err%s", err)
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
	entry.Creator = userId
	entry.AppId = appId
	tagId, err := handler.AddTag(ctx, appId, entry)
	if err == nil {
		result = &httpDto.CreateTagRpn{
			TagId: tagId,
		}
	}
}

func UpdateTag(c *gin.Context) {
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
	tagId := c.Param("tagId")
	entry := &entries.UpdateTagEntry{}
	if err := c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err = handler.UpdateTag(ctx, tagId, entry)
}

func DeleteTag(c *gin.Context) {
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
	tagId := c.Param("tagId")
	err = handler.DeleteTag(ctx, tagId)
}

func QueryTag(c *gin.Context) {
	var (
		err    error
		result []*dao.Tag_DB
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
	classId := c.Param("classId")
	entry := &entries.QueryTagEntry{
		Owner: c.Query("owner"),
		Type:  c.Query("type"),
	}
	result, err = handler.QueryTag(ctx, appId, classId, entry)
}

func GetTag(c *gin.Context) {
	var (
		err    error
		result *dao.Tag_DB
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
	tagId := c.Param("tagId")
	result, err = handler.GetTag(ctx, tagId)
}

// QueryTagList 查询标签列表
func QueryTagList(c *gin.Context) {
	var (
		err    error
		result []*dao.Tag_DB
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
	entry := &httpDto.QueryTagReq{
		Owner: c.Query("owner"),
		Type:  c.Query("type"),
	}
	result, err = handler.QueryTagList(ctx, appId, entry)
}
