package main

import (
	"assetLibary/configs"
	"assetLibary/dao"
	"assetLibary/restful"
	"assetLibary/services"
	"gitlab.zixel.cn/go/framework"
)

func main() {
	dao.Init()
	services.Init()

	for k, v := range configs.ErrorCodes {
		framework.SetErrorTips(k, v)
	}
	framework.LoadRoute(restful.InitRoute)

	s := framework.GetGrpcServer()

	services.RegisterAssetLibrarySerivceServer(s, restful.NewService())
	framework.Run()
}
