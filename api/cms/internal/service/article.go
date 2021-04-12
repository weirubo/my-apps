package service

import "my-apps/api/cms/internal/model"

type ArticleReq struct {
	ID           uint   `json:"id" form:"id"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Content      string `json:"content" binding:"required"`
	CategoryID   uint8  `json:"category_id" binding:"required"`
	CategoryName string `json:"category_name" binding:"required"`
	TagID        uint8  `json:"tag_id" binding:"required"`
	TagName      string `json:"tag_name" binding:"required"`
	CommentCount uint   `json:"comment_count" binding:"required"`
}

func (s service) AddArticle(param *ArticleReq) error {
	return s.dao.CreateArticle(param.Title, param.Description, param.Content, param.CategoryName, param.TagName, param.CategoryID, param.TagID, param.CommentCount)
}

func (s service) GetArticles(param *Page) ([]model.Article, error) {
	return s.dao.ReadArticles(param.PageNumber, param.PageSize)
}

func (s service) UpdateArticle(param *ArticleReq) error {
	return s.dao.UpdateArticle(param.ID, param.CommentCount, param.Title, param.Description, param.Content, param.CategoryName, param.TagName, param.CategoryID, param.TagID)
}

func (s service) DeleteArticle(param *ArticleReq) error {
	return s.dao.DeleteArticle(param.ID)
}

func (s service) GetArticle(param *ArticleReq) (model.ArticleView, error) {
	return s.dao.ReadArticle(param.ID)
}
