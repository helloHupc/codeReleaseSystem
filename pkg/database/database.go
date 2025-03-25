package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println("gorm open err", err.Error())
	}

	// 将 *gorm.DB 赋值给 SQLDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println("sqldb err", err.Error())
	}
}
