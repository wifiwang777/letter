package controller

import (
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	err = service.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg("register success"))
}

func Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	jwt, err := service.UserService.Login(&user)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg(jwt))
}
