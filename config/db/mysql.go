package db

import (
	"github.com/aisuosuo/letter/api/models"
	"github.com/aisuosuo/letter/config"
	"github.com/aisuosuo/letter/config/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := config.GlobalConfig.Mysql.Dsn()
	var err error

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	err = db.AutoMigrate(&models.User{}, &models.Messages{}, &models.UserFriend{})
	if err != nil {
		log.Logger.Error(err)
	}
	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(config.GlobalConfig.Mysql.MaxOpenConns) //设置数据库连接池最大连接数20
	sqlDB.SetMaxIdleConns(config.GlobalConfig.Mysql.MaxIdleConns) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于10，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return db
}
