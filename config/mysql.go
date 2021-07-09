package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var MysqlClient *gorm.DB

func InitMysql() {
	dsn := "root:123456@(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}