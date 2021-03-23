package dao

import "gorm.io/gorm"

// 数据操作层

type Dao struct {
	dbEngine *gorm.DB
}

// 实例化
func New(dbEngine *gorm.DB) *Dao {
	return &Dao{dbEngine: dbEngine}
}
