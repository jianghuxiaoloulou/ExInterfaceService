package routers

import (
	_ "WowjoyProject/ExInterfaceService/docs"
	v1 "WowjoyProject/ExInterfaceService/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册中间件
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 获取读卡信息
		apiv1.GET("/ReadCardInfo", v1.ReadCardInfo)
	}
	return r
}
