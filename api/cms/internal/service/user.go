package service

import (
	"my-apps/api/cms/internal/model"
)

type UserReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (s service) AddUser(param *UserReq) error {
	return s.dao.CreateUser(param.UserName, param.Email, param.PassWord)
}

func (s service) GetUsers(param *Page) ([]model.User, error) {
	return s.dao.ReadUsers(param.PageNumber, param.PageSize)
}
