package model

import (
	"fmt"
	"gorm.io/gorm"
)

// 映射 articles 表
type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(50) not null;commit:标题"`
	Description  string `gorm:"type:varchar(360) not null;commit:描述"`
	CategoryID   uint8  `gorm:"type:smallint not null;commit:分类ID"`
	CategoryName string `gorm:"type:varchar(20) not null;commit:分类名称"`
	CommentCount uint   `gorm:"type:int(10) not null;commit:评论总数"`
}

// 表操作
func (a Article) Create(db *gorm.DB) error {
	err := db.AutoMigrate(a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&a).Error
}
