package api

import (
	"github.com/aisuosuo/letter/api/controller"
	"github.com/aisuosuo/letter/config"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func init() {
	if config.GlobalConfig.Run.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	HttpServer = gin.Default()
	HttpServer.Use(Cors())

	userGroup := HttpServer.Group("user")
	userGroup.POST("register", controller.Register)
	userGroup.POST("login", controller.Login)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		c.Next()
	}
}
