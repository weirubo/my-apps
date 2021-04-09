package dao

import (
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

func (d *Dao) CreateComment(content string, articleID, commentID uint) error {
	comment := model.Comment{
		Model:     gorm.Model{},
		Content:   content,
		ArticleID: articleID,
		CommentID: commentID,
	}
	return comment.Create(d.dbEngine)
}
