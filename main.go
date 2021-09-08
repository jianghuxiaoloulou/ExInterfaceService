package main

import (
	"WowjoyProject/ExInterfaceService/global"
	"WowjoyProject/ExInterfaceService/internal/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title 对外接口服务
// @version 1.0.0.1
// @description 对接第三方服务接口
// @termsOfService https://github.com/jianghuxiaoloulou/ExInterfaceService.git
func main() {
	global.Logger.Info("*******开始运行对外接口服务********")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	ser.ListenAndServe()
}
