package service

import (
	"my-apps/api/cms/common"
	"my-apps/api/cms/internal/dao"
)

type service struct {
	dao *dao.Dao
}

type Page struct {
	PageNumber int `form:"page_number"`
	PageSize   int `form:"page_size"`
}

// 实例化
func New() service {
	svc := service{
		dao: dao.New(common.DBEngine),
	}
	return svc
}
