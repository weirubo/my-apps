package dao

import (
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

// 数据操作层

func (d *Dao) Create(username, email, password string) error {
	user := model.User{
		Model:    gorm.Model{},
		Username: username,
		Email:    email,
		Password: password,
	}
	return user.Create(d.dbEngine)
}
