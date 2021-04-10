package service

import (
	"my-apps/api/cms/internal/model"
)

type CommentReq struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	ArticleID uint   `json:"article_id"`
	CommentID uint   `json:"comment_id"`
}

func (s service) AddComment(param *CommentReq) error {
	return s.dao.CreateComment(param.Content, param.ArticleID, param.CommentID)
}

func (s service) GetComments(param *Page) ([]model.Comment, error) {
	return s.dao.ReadComments(param.PageNumber, param.PageSize)
}

func (s service) UpdateComment(param *CommentReq) error {
	return s.dao.UpdateComment(param.ID, param.Content)
}

func (s service) DeleteComment(param *CommentReq) error {
	return s.dao.DeleteComment(param.ID)
}
