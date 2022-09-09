package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/aisuosuo/letter/api/jwt"
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

func GetUidFromToken(tokenStr string) (int, error) {
	claims, err := jwt.Verify(tokenStr)
	if err != nil {
		return 0, err
	}
	var uid int
	switch claims["uid"].(type) {
	case float64:
		uid = int(claims["uid"].(float64))
	case int:
		uid = claims["uid"].(int)
	case string:
		uid, _ = strconv.Atoi(claims["uid"].(string))
	default:
		return 0, errors.New("invalid token")
	}
	return uid, nil
}
