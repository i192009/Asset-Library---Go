package router

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/handler"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func AddInstanceScopeId(c *gin.Context) {
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
	entry := &entries.CreateInstanceScopeIdEntry{}
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
	result, err = handler.AddInstanceScopeId(ctx, entry)
}

func UpdateInstanceScopeId(c *gin.Context) {
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
	entry := &entries.UpdateInstanceScopeIdEntry{}
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
	err = handler.UpdateInstanceScopeId(ctx, entry)
}

func DeleteInstanceScopeId(c *gin.Context) {
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
	instanceId := c.Param("instanceId")
	err = handler.DeleteInstanceScopeId(ctx, instanceId)
}

func GetInstanceScopeId(c *gin.Context) {
	var (
		err    error
		result *dao.InstanceScopeId_DB
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
	instanceId := c.Param("instanceId")
	result, err = handler.GetInstanceScopeId(ctx, instanceId)
}

func QueryInstanceScopeIdList(c *gin.Context) {
	var (
		err    error
		result []*dao.InstanceScopeId_DB
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

	result, err = handler.QueryInstanceScopeIdList(ctx)
}
