package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;index:idx_username_email;commit:用户名"`
	Email    string `gorm:"type:varchar(20);not null;unique;index:idx_username_email;commit:邮箱"`
	Password string `gorm:"type:varchar(20);not null;commit:密码" json:"-"`
}

func (u User) Create(db *gorm.DB) error {
	err := db.AutoMigrate(u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&u).Error
}

func (u User) ReadByPage(db *gorm.DB, pageNumber, pageSize int) ([]User, error) {
	var users []User
	pageOffset := (pageNumber - 1) * pageSize
	tx := db.Select("id", "username", "email", "created_at", "updated_at").Order("id desc").Limit(pageSize).Offset(pageOffset).Find(&users)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) UpdateByID(db *gorm.DB) error {
	tx := db.Where("id", u.ID).Updates(&u)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
