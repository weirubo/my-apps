package service

import (
	"my-apps/api/cms/internal/model"
)

type UserReq struct {
	ID       uint   `json:"id"`
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

func (s service) AddUser(param *UserReq) error {
	return s.dao.CreateUser(param.UserName, param.Email, param.PassWord)
}

func (s service) GetUsers(param *Page) ([]model.User, error) {
	return s.dao.ReadUsers(param.PageNumber, param.PageSize)
}

func (s service) UpdateUser(param *UserReq) error {
	return s.dao.UpdateUser(param.ID, param.UserName, param.Email, param.PassWord)
}

func (s service) DeleteUser(param *UserReq) error {
	return s.dao.DeleteUser(param.ID)
}
