package model

import (
	"fmt"
	"gorm.io/gorm"
)

// 映射 tags 表
type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(20) not null;commit:标签名称"`
	Count   uint   `gorm:"type:int(10) not null;commit:文章总数"`
}

// 表操作
func (t Tag) Create(db *gorm.DB) error {
	err := db.AutoMigrate(t)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&t).Error
}
