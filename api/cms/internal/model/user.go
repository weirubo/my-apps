package model

import (
	"fmt"
	"gorm.io/gorm"
)

// 映射 user 表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;index:idx_username_email;commit:用户名"`
	Email    string `gorm:"type:varchar(20);not null;unique;index:idx_username_email;commit:邮箱"`
	Password string `gorm:"type:varchar(20);not null;commit:密码"`
}

// 表操作
func (u User) Create(db *gorm.DB) error {
	err := db.AutoMigrate(u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(u).Error
}
