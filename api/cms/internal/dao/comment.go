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
	// TODO::更新文章表的评论总数
	return comment.Create(d.dbEngine)
}

func (d *Dao) ReadComments(pageNumber, pageSize int) ([]model.Comment, error) {
	comment := model.Comment{}
	return comment.ReadByPage(d.dbEngine, pageNumber, pageSize)
}

func (d *Dao) UpdateComment(id uint, content string) error {
	comment := model.Comment{
		Model: gorm.Model{
			ID: id,
		},
		Content: content,
	}
	return comment.UpdateByID(d.dbEngine)
}

func (d *Dao) DeleteComment(id uint) error {
	comment := model.Comment{
		Model: gorm.Model{
			ID: id,
		},
	}
	return comment.DeleteByID(d.dbEngine)
}
