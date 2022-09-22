package controller

import (
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

func GetFriends(c *gin.Context) {
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	friends := service.UserService.GetFriends(uid)
	c.JSON(http.StatusOK, service.SuccessMsg(friends))
}

func SearchUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusOK, service.FailMsg("invalid username"))
		return
	}
	user, err := service.UserService.GetUser(name)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg(user))
}

func AddFriend(c *gin.Context) {
	var friend models.User
	err := c.ShouldBindJSON(&friend)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	err = service.UserService.AddFriend(uid, friend.UID)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "Duplicate entry") {
			errMsg = "已添加该好友"
		}
		c.JSON(http.StatusOK, service.FailMsg(errMsg))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg("add friend success"))
}

func DeleteFriend(c *gin.Context) {
	var friend models.User
	err := c.ShouldBindJSON(&friend)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	err = service.UserService.DeleteFriend(uid, friend.UID)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg("delete friend success"))
}

func GetMessages(c *gin.Context) {
	friend := c.Query("uid")
	friendUid, err := strconv.Atoi(friend)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	messages := service.UserService.GetMessages(uid, uint(friendUid))
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg(messages))
}
