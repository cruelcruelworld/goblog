package model

import (
	"fmt"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	var err error

	config := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
			config.GetString("database.mysql.username"),
			config.GetString("database.mysql.password"),
			config.GetString("database.mysql.host"),
			config.GetString("database.mysql.port"),
			config.GetString("database.mysql.database"),
			config.GetString("database.mysql.charset"),
			),
	})

	DB, err = gorm.Open(config, &gorm.Config{
			Logger: logger2.Default.LogMode(logger2.Warn),
	})

	logger.LogError(err)

	return DB
}
