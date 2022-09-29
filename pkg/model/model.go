package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:root@tcp(192.168.0.104:3307)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	DB, err := gorm.Open(config, &gorm.Config{})

	logger.LogError(err)

	return DB
}
