package service

import "my-apps/api/cms/internal/model"

type ArticleReq struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
	CategoryID   uint8  `json:"category_id"`
	CategoryName string `json:"category_name"`
	TagID        uint8  `json:"tag_id"`
	TagName      string `json:"tag_name"`
	CommentCount uint   `json:"comment_count"`
}

func (s service) AddArticle(param *ArticleReq) error {
	return s.dao.CreateArticle(param.Title, param.Description, param.Content, param.CategoryName, param.TagName, param.CategoryID, param.TagID, param.CommentCount)
}

func (s service) GetArticles(param *Page) ([]model.Article, error) {
	return s.dao.ReadArticles(param.PageNumber, param.PageSize)
}
