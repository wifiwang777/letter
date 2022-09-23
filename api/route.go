package api

import (
	"github.com/aisuosuo/letter/api/controller"
	"github.com/aisuosuo/letter/api/jwt"
	"github.com/aisuosuo/letter/api/service"
	"github.com/aisuosuo/letter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var HttpServer *gin.Engine

func init() {
	if config.GlobalConfig.Run.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	HttpServer = gin.Default()
	HttpServer.Use(Cors())

	//base api
	HttpServer.POST("register", controller.Register)
	HttpServer.POST("login", controller.Login)

	privateGroup := HttpServer.Group("")
	privateGroup.Use(JWTAuth)

	userGroup := privateGroup.Group("user")
	userGroup.GET("getFriends", controller.GetFriends)
	userGroup.GET("info", controller.UserInfo)
	userGroup.GET("searchUser", controller.SearchUser)
	userGroup.POST("addFriend", controller.AddFriend)
	userGroup.POST("deleteFriend", controller.DeleteFriend)

	messageGroup := privateGroup.Group("messages")
	messageGroup.GET("", controller.GetMessages)
}

func Cors() gin.HandlerFunc {
	corConf := cors.DefaultConfig()
	corConf.AllowAllOrigins = true
	corConf.AllowHeaders = append(corConf.AllowHeaders, "X-Token")
	return cors.New(corConf)
}

func JWTAuth(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		c.JSON(http.StatusOK, service.FailMsg("非法访问"))
		c.Abort()
		return
	}
	claims, err := jwt.Verify(token)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
