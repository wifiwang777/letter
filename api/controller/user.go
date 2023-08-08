package controller

import (
	"fmt"
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/api/service"
	"github.com/aisuosuo/letter/core/pb"
	"github.com/aisuosuo/letter/core/ws"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
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
	user, err := service.UserService.SearchUser(name)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg(user))
}

func UserInfo(c *gin.Context) {
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	user, err := service.UserService.UserInfo(uid)
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

	if service.UserService.IsFriend(uid, friend.UID) {
		c.JSON(http.StatusOK, service.FailMsg("this user has been your friend"))
		return
	}

	user, err := service.UserService.UserInfo(uid)
	if err != nil {
		errMsg := err.Error()
		c.JSON(http.StatusOK, service.FailMsg(errMsg))
		return
	}

	message := new(pb.Message)
	message.From = uint32(uid)
	message.To = uint32(friend.UID)
	message.Type = pb.MessageType_Notify
	message.Content = fmt.Sprintf("accept new friend reqeust from %s ?", user.Name)

	bytes, err := proto.Marshal(message)
	if err != nil {
		errMsg := err.Error()
		c.JSON(http.StatusOK, service.FailMsg(errMsg))
		return
	}
	ws.WsServer.ReceiveMessage(bytes)
	c.JSON(http.StatusOK, service.SuccessMsg("sent new friend request success"))
	return
}

func AcceptFriend(c *gin.Context) {
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}

	var friend models.User
	err := c.ShouldBindJSON(&friend)
	if err != nil {
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}

	err = service.UserService.AcceptFriend(uid, friend.UID)
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
	c.JSON(http.StatusOK, service.SuccessMsg(service.ConvertMessageTime(messages)))
}
