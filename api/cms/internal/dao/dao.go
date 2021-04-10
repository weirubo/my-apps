package dao

import "gorm.io/gorm"

type Dao struct {
	dbEngine *gorm.DB
}

func New(dbEngine *gorm.DB) *Dao {
	return &Dao{dbEngine: dbEngine}
}
