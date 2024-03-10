package router

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful/contants"
	"assetLibary/restful/entries"
	"assetLibary/restful/handler"
	"assetLibary/restful/httpDto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func AddClass(c *gin.Context) {
	var (
		err    error
		result *httpDto.CreateClassRpn
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

	entry := &entries.CreateClassEntry{}
	if err = c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	if err = ValidateJSON(entry); err != nil {
		err = CreateError(codes.InvalidArgument, 50002, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	//
	entry.Creator = userId
	classId, err := handler.AddClass(ctx, appId, entry)
	if err == nil {
		result = &httpDto.CreatCheClassRpn{
			ClassId: classId,
		}
	}
}

func UpdateClass(c *gin.Context) {
	var (
		err    error
		result string
	)
	ctx, cancel := context.WithTimeout(context.Background(), configs.Timeout)
	defer cancel()

	defer func() {
		if err == nil {
			fmt.Println("CT-3")
			c.JSON(200, result)
		} else {
			SendErrorToJson(c, err)
		}
	}()
	classId := c.Param("classId")
	//appId := c.GetHeader(contants.ZixelHeader_AppId)
	entry := &entries.UpdateClassEntry{}
	if err = c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err = handler.UpdateClass(ctx, classId, entry)
}

func DeleteClass(c *gin.Context) {
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
	classId := c.Param("classId")
	//appId := c.GetHeader(contants.ZixelHeader_AppId)
	err = handler.DeleteClass(ctx, classId)
}

func QueryClass(c *gin.Context) {
	var (
		err    error
		result []*dao.Class_DB
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
	result, err = handler.QueryClass(ctx, appId)
}

func GetClass(c *gin.Context) {
	var (
		err    error
		result *dao.Class_DB
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
	classId := c.Param("classId")
	result, err = handler.GetClass(ctx, classId)
}
