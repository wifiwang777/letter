package gorm

import (
	"fmt"
	"github.com/aisuosuo/letter/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", "root", "123456", "127.0.0.1", "3306", "letter", "charset=utf8mb4&parseTime=True&loc=Local")
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	UserModel := models.UserMgr(db)
	var count int64
	UserModel.Where("name", "wf").Count(&count)
	t.Log(count)

}
