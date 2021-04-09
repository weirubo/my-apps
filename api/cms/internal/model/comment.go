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

func (c Comment) ReadByPage(db *gorm.DB, pageNumber, pageSize int) ([]Comment, error) {
	var comments []Comment
	pageOffset := (pageNumber - 1) * pageSize
	tx := db.Select("id", "content", "article_id", "comment_id", "created_at", "updated_at").Order("id desc").Limit(pageSize).Offset(pageOffset).Find(&comments)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return comments, nil
}
