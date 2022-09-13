package api

import (
	"github.com/aisuosuo/letter/api/controller"
	"github.com/aisuosuo/letter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func init() {
	if config.GlobalConfig.Run.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	HttpServer = gin.Default()
	HttpServer.Use(cors.Default())

	userGroup := HttpServer.Group("user")
	userGroup.POST("register", controller.Register)
	userGroup.POST("login", controller.Login)
}
