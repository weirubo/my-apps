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

func (d *Dao) ReadComments(pageNumber, pageSize int) ([]model.Comment, error) {
	comment := model.Comment{}
	return comment.ReadByPage(d.dbEngine, pageNumber, pageSize)
}
