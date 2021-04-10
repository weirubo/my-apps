package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string `gorm:"type:varchar(50) not null;commit:标题"`
	Description  string `gorm:"type:varchar(360) not null;commit:描述"`
	CategoryID   uint8  `gorm:"type:smallint not null;commit:分类ID"`
	CategoryName string `gorm:"type:varchar(20) not null;commit:分类名称"`
	CommentCount uint   `gorm:"type:int(10) not null;commit:评论总数"`
}

func (a Article) Create(db *gorm.DB) error {
	err := db.AutoMigrate(a)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&a).Error
}

func (a Article) ReadByPage(db *gorm.DB, pageNumber, pageSize int) ([]Article, error) {
	var articles []Article
	pageOffset := (pageNumber - 1) * pageSize
	tx := db.Select("id", "title", "description", "category_id", "category_name", "comment_count", "created_at", "updated_at").Limit(pageSize).Offset(pageOffset).Find(&articles)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return articles, nil
}
