package router

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful/entries"
	"assetLibary/restful/handler"
	"assetLibary/restful/httpDto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func AddConfig(c *gin.Context) {
	var (
		err    error
		result *httpDto.CreateConfigDto
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
	entry := &entries.CreateConfigEntry{}
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
	configId, err := handler.AddConfig(ctx, entry)
	if err == nil {
		result = &httpDto.CreateConfigDto{
			ConfigId: configId,
		}
	}
}

func UpdateConfig(c *gin.Context) {
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
	configId := c.Param("configId")
	entry := &entries.UpdateConfigEntry{}
	if err = c.ShouldBindJSON(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 50001, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	err = handler.UpdateConfig(ctx, configId, entry)
}

func DeleteConfig(c *gin.Context) {
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
	configId := c.Param("configId")
	err = handler.DeleteConfig(ctx, configId)
}

func GetConfig(c *gin.Context) {
	var (
		err    error
		result *dao.Config_DB
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
	configId := c.Param("configId")
	result, err = handler.GetConfig(ctx, configId)
}

func QueryConfig(c *gin.Context) {
	var (
		err    error
		result []*dao.Config_DB
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
	result, err = handler.QueryConfig(ctx)
}

func QueryConfigByType(c *gin.Context) {
	var (
		err    error
		result *dao.Config_DB
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
	entry := &entries.QueryConfigEntry{}
	if err := c.ShouldBindUri(entry); err != nil {
		log.Error(err.Error())
		err = CreateError(codes.InvalidArgument, 40001, map[string]interface{}{
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
	result, err = handler.QueryConfigByType(ctx, entry)
}
