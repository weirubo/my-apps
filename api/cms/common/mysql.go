package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建数据库

func NewDBEngine() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
		DSN: "root:root@tcp(127.0.0.1:3306)/cms?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
