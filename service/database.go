package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_URL = "root:123456@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local"

// 连接数据库
func GetDb() (db *gorm.DB, err error) {

	return gorm.Open(mysql.Open(DB_URL), &gorm.Config{
		// 日志级别
		Logger: logger.Default.LogMode(logger.Info),
	})

}
