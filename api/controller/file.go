package controller

import (
	"fmt"
	"github.com/aisuosuo/letter/api/service"
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/config/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GetFile(c *gin.Context) {
	fileName := c.Param("fileName")
	filePath := fmt.Sprintf("%s/%s", config.GlobalConfig.FileConfig.FilePath, fileName)
	data, _ := os.ReadFile(filePath)
	c.Writer.Write(data)
}

func SaveFile(c *gin.Context) {
	uid := service.GetUserId(c)
	if uid == 0 {
		c.JSON(http.StatusOK, service.FailMsg("invalid token"))
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Logger.Error(err)
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	fileName := file.Filename
	log.Logger.Debugf("fileName:%s", fileName)
	c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", config.GlobalConfig.FileConfig.FilePath, fileName))

	err = service.UserService.UpdateAvatar(uid, fileName)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(http.StatusOK, service.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.SuccessMsg("update avatar success"))
}
