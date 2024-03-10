package restful

import (
	"assetLibary/restful/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(app *gin.RouterGroup) {
	backend := app.Group("/assetManage/v1/backend")

	backendAsset := backend.Group("/asset")
	backendAsset.POST("/add", router.AddAsset)
	backendAsset.POST("/:permissionType/query", router.QueryAsset)
	backendAsset.POST("/assetUploaded/:permissionType/:assetId", router.AssetUploaded)
	backendAsset.PUT("/update/:permissionType/:assetId", router.UpdateAsset)
	backendAsset.PUT("/update/:permissionType/:assetId/thumbnail", router.UpdateAssetThumbnail)
	backendAsset.DELETE("/delete/:permissionType/:assetId", router.DeleteAsset)
	backendAsset.GET("/get/:permissionType/:assetId", router.GetAsset)
	backendAsset.POST("/copy/:permissionType/:assetId", router.UseAsset)
	backendAsset.POST("/checkPermission", router.CheckPermission)

	backendClass := backend.Group("/class")
	backendClass.POST("/add", router.AddClass)
	backendClass.PUT("/update/:classId", router.UpdateClass)
	backendClass.DELETE("/delete/:classId", router.DeleteClass)
	backendClass.GET("/query", router.QueryClass)
	backendClass.GET("/get/:classId", router.GetClass)

	backendTag := backend.Group("/tag")
	backendTag.POST("/add", router.AddTag)
	backendTag.PUT("/update/:tagId", router.UpdateTag)
	backendTag.DELETE("/delete/:tagId", router.DeleteTag)
	backendTag.GET("/query/:classId", router.QueryTag)
	backendTag.GET("/get/:tagId", router.GetTag)
	backendTag.GET("/query/list", router.QueryTagList)

	backendConfig := backend.Group("/config")
	backendConfig.POST("/add", router.AddConfig)
	backendConfig.PUT("/update/:configId", router.UpdateConfig)
	backendConfig.GET("/get/:configId", router.GetConfig)
	backendConfig.GET("/query", router.QueryConfig)
	backendConfig.GET("/query/:type/:owner", router.QueryConfigByType)
	backendConfig.DELETE("/delete/:configId", router.DeleteConfig)

	backendInstance := backend.Group("/instanceScopeId")
	backendInstance.POST("/add", router.AddInstanceScopeId)
	backendInstance.PUT("/update", router.UpdateInstanceScopeId)
	backendInstance.GET("/get/:instanceScopeId", router.GetInstanceScopeId)
	backendInstance.GET("/query", router.QueryInstanceScopeIdList)
	backendInstance.DELETE("/delete/:instanceScopeId", router.DeleteInstanceScopeId)

	app.GET("/health", health)
	app.GET("/available", available)
}

func health(c *gin.Context) {
	//log.Info("health check")
	c.JSON(http.StatusOK, "ok")
}

func available(c *gin.Context) {
	//log.Info("status check")
	c.JSON(http.StatusOK, "ok")
}
