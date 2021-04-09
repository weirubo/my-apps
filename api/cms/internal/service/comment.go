package service

type Comment struct {
	Content   string `json:"content"`
	ArticleID uint   `json:"article_id"`
	CommentID uint   `json:"comment_id"`
}

func (s service) AddComment(param *Comment) error {
	return s.dao.CreateComment(param.Content, param.ArticleID, param.CommentID)
}
