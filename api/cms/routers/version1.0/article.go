package version1_0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
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
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.AddArticle(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (a Article) List(c *gin.Context) {
	param := &service.Page{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	articles, err := svc.GetArticles(param)
	if err != nil {
		fmt.Println(err)
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
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.UpdateArticle(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
