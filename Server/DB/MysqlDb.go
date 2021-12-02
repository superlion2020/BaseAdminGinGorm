package DB

import (
	"Server/Models/DbModel"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MysqlDB *gorm.DB

func init() {
	fmt.Println("mysqldb初始化开始……")
	dsn := "root:!QAZxsw2@tcp(127.0.0.1:23001)/go_gin_api?charset=utf8&parseTime=True&loc=Local"
	//dsn:= "user=root password=!QAZxsw2 dbname=go_gin_api port=23001 sslmode=disable TimeZone=Asia/Shanghai"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("mysql 连接错误")
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	gormDB.AutoMigrate(&DbModel.User{})
	MysqlDB = gormDB
	fmt.Println("mysqldb初始化结束")
}
