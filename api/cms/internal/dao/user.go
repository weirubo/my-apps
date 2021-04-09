package dao

import (
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

// 数据操作层
func (d *Dao) CreateUser(username, email, password string) error {
	user := model.User{
		Model:    gorm.Model{},
		Username: username,
		Email:    email,
		Password: password,
	}
	return user.Create(d.dbEngine)
}

func (d *Dao) ReadUsers(pageNumber, pageSize int) ([]model.User, error) {
	user := model.User{}
	return user.ReadByPage(d.dbEngine, pageNumber, pageSize)
}
