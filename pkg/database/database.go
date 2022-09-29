package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"goblog/pkg/logger"
	"time"
)

var DB *sql.DB

func Initialize()  {
	initDB()
	createTables()
}

func initDB()  {
	var err error
	config := mysql.Config{
		User: "root",
		Passwd: "root",
		Addr: "192.168.0.104:3307",
		Net: "tcp",
		DBName: "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    body longtext COLLATE utf8mb4_general_ci
); `

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}