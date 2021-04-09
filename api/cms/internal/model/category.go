package model

import (
	"fmt"
	"gorm.io/gorm"
)

// 映射 categories 表
type Category struct {
	gorm.Model
	CategoryName string `gorm:"type:varchar(20);not null;commit:分类名称"`
	Count        uint   `gorm:"type:int(10);not null;commit:文章总数"`
}

// 表操作
func (c Category) Create(db *gorm.DB) error {
	err := db.AutoMigrate(c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&c).Error
}
