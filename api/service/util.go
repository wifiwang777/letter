package service

import (
	"crypto/md5"
	"fmt"
	"github.com/aisuosuo/letter/config/log"
	"github.com/gin-gonic/gin"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"strconv"
)

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data interface{}) *ResponseMsg {
	msg := &ResponseMsg{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailMsg(msg string) *ResponseMsg {
	msgObj := &ResponseMsg{
		Code: -1,
		Msg:  msg,
	}
	return msgObj
}

func PasswordEncrypt(original string) string {
	encryptPassword := md5.Sum([]byte(original))

	return fmt.Sprintf("%x", encryptPassword)
}

func GetUserId(c *gin.Context) uint {
	if value, ok := c.Get("claims"); !ok {
		return 0
	} else {
		claims := value.(jwt2.MapClaims)
		var uid int
		switch claims["uid"].(type) {
		case float64:
			uid = int(claims["uid"].(float64))
		case int:
			uid = claims["uid"].(int)
		case string:
			uid, _ = strconv.Atoi(claims["uid"].(string))
		default:
			log.Logger.Error("invalid uid format")
			return 0
		}
		return uint(uid)
	}
}
