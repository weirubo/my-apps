package version1_0

import (
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"my-apps/api/cms/pkg"
	"net/http"
)

type Article struct {
}

type ArticleRes struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	CategoryID   uint8  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CommentCount uint   `json:"comment_count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (a Article) Add(c *gin.Context) {
	param := &service.ArticleReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.AddArticle(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20401,
			"msg":  "add article failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (a Article) List(c *gin.Context) {
	param := &service.Page{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	articles, err := svc.GetArticles(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20402,
			"msg":  "get articles failed",
		})
		return
	}
	var article ArticleRes
	var articleList []ArticleRes
	for _, v := range articles {
		article.ID = v.ID
		article.Title = v.Title
		article.Description = v.Description
		article.CategoryID = v.CategoryID
		article.CategoryName = v.CategoryName
		article.CommentCount = v.CommentCount
		article.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		article.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		articleList = append(articleList, article)
	}
	data := gin.H{
		"articles": articleList,
	}
	c.JSON(http.StatusOK, data)
}

func (a Article) Edit(c *gin.Context) {
	param := &service.ArticleReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.UpdateArticle(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20403,
			"msg":  "edit article failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (a Article) Remove(c *gin.Context) {
	param := &service.ArticleReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.DeleteArticle(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20404,
			"msg":  "delete article failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (a Article) View(c *gin.Context) {
	param := &service.ArticleReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	articleView, err := svc.GetArticle(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20405,
			"msg":  "get article failed",
		})
		return
	}
	articleData := ArticleRes{
		ID:           articleView.ID,
		Title:        articleView.Title,
		CategoryID:   articleView.CategoryID,
		CategoryName: articleView.CategoryName,
		CommentCount: articleView.CommentCount,
		CreatedAt:    articleView.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    articleView.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusOK, articleData)
}
