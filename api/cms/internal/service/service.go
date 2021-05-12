package service

import (
	common2 "my-apps/api/cms/common"
	"my-apps/api/cms/internal/dao"
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
		dao: dao.New(common2.DBEngine),
	}
	return svc
}
