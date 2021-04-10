package model

import (
	"fmt"
	"gorm.io/gorm"
)

type ArticleView struct {
	gorm.Model
	Title        string `gorm:"type:varchar(50) not null;commit:标题"`
	Content      string `gorm:"type:text not null;commit:内容"`
	CategoryID   uint8  `gorm:"type:smallint not null;commit:分类ID"`
	CategoryName string `gorm:"type:varchar(20) not null;commit:分类名称"`
	TagID        uint8  `gorm:"type:smallint not null;commit:标签ID"`
	TagName      string `gorm:"type:varchar(20) not null;commit:标签名称"`
	CommentCount uint   `gorm:"type:int(10) not null;commit:评论总数"`
}

func (av ArticleView) Create(db *gorm.DB) error {
	err := db.AutoMigrate(av)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&av).Error
}
