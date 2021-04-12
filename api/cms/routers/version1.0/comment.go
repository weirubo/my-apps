package version1_0

import (
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"my-apps/api/cms/pkg"
	"net/http"
)

type Comment struct {
}

type CommentRes struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserID    uint   `json:"uid"`
	UserName  string `json:"username"`
	ArticleID uint   `json:"article_id"`
	CommentID uint   `json:"comment_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (com Comment) Add(c *gin.Context) {
	param := &service.CommentReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.AddComment(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20501,
			"msg":  "add comment failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (com Comment) List(c *gin.Context) {
	param := &service.Page{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	comments, err := svc.GetComments(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20502,
			"msg":  "get comments failed",
		})
		return
	}
	var comment CommentRes
	var commentList []CommentRes
	for _, v := range comments {
		comment.ID = v.ID
		comment.Content = v.Content
		comment.UserID = v.UserID
		comment.UserName = v.UserName
		comment.ArticleID = v.ArticleID
		comment.CommentID = v.CommentID
		comment.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		comment.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		commentList = append(commentList, comment)
	}
	data := gin.H{
		"comments": commentList,
	}
	c.JSON(http.StatusOK, data)
}

func (com Comment) Edit(c *gin.Context) {
	param := &service.CommentReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.UpdateComment(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20503,
			"msg":  "edit comment failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (com Comment) Remove(c *gin.Context) {
	param := &service.CommentReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.DeleteComment(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20504,
			"msg":  "delete comment failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
