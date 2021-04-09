package service

import (
	"my-apps/api/cms/common"
	"my-apps/api/cms/internal/dao"
)

type service struct {
	dao *dao.Dao
}

type Page struct {
	PageSize   int `form:"page_size"`
	PageOffset int `form:"page_offset"`
}

// 实例化
func New() service {
	svc := service{
		dao: dao.New(common.DBEngine),
	}
	return svc
}
