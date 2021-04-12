package service

import (
	"my-apps/api/cms/internal/dao"
	"my-apps/api/cms/pkg"
)

type service struct {
	dao *dao.Dao
}

type Page struct {
	PageNumber int `form:"page_number"`
	PageSize   int `form:"page_size"`
}

func New() service {
	svc := service{
		dao: dao.New(pkg.DBEngine),
	}
	return svc
}
