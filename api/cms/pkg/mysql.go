package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 创建数据库

func NewDBEngine(databaseConfig *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.DBName, databaseConfig.Charset, databaseConfig.ParseTime)
	db, err := gorm.Open(mysql.New(mysql.Config{
		// gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local
		// DSN: "root:root@tcp(127.0.0.1:3306)/cms?charset=utf8&parseTime=True&loc=Local",
		DSN: dsn,
	}), &gorm.Config{SkipDefaultTransaction: false})
	if err != nil {
		return nil, err
	}
	return db, nil
}
