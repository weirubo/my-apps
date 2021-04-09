package model

import (
	"fmt"
	"gorm.io/gorm"
)

// 映射 comments 表
type Comment struct {
	gorm.Model
	Content   string `gorm:"type:varchar(360) not null;commit:评论内容"`
	ArticleID uint   `gorm:"type:int(11) not null;commit:文章ID"`
	CommentID uint   `gorm:"type:int(11) not null;commit:回复评论ID"`
}

// 表操作
func (c Comment) Create(db *gorm.DB) error {
	err := db.AutoMigrate(c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&c).Error
}
