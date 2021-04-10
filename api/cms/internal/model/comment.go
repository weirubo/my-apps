package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:varchar(360) not null;commit:评论内容"`
	UserID    uint   `gorm:"type:int(11) not null;commit:用户ID"`
	UserName  string `gorm:"type:varchar(20) not null;commit:用户名"`
	ArticleID uint   `gorm:"type:int(11) not null;commit:文章ID"`
	CommentID uint   `gorm:"type:int(11) not null;commit:回复评论ID"`
}

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
	tx := db.Select("id", "content", "user_id", "user_name", "article_id", "comment_id", "created_at", "updated_at").Order("id desc").Limit(pageSize).Offset(pageOffset).Find(&comments)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return comments, nil
}

func (c Comment) UpdateByID(db *gorm.DB) error {
	tx := db.Where("id", c.ID).Updates(&c)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (c Comment) DeleteByID(db *gorm.DB) error {
	tx := db.Delete(&c, c.ID)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
