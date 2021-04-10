package routers

import (
	"github.com/gin-gonic/gin"
)
import version "my-apps/api/cms/routers/version1.0"

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	// 不使用默认中间件
	r := gin.New()
	user := version.User{}
	category := version.Category{}
	tag := version.Tag{}
	article := version.Article{}
	comment := version.Comment{}
	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.POST("/user/add", user.Add)
		v1.GET("/user/list", user.List)
		v1.POST("/user/edit", user.Edit)
		v1.POST("/category/add", category.Add)
		v1.GET("/category/list", category.List)
		v1.POST("/category/edit", category.Edit)
		v1.POST("/tag/add", tag.Add)
		v1.GET("/tag/list", tag.List)
		v1.POST("/tag/edit", tag.Edit)
		v1.POST("/article/add", article.Add)
		v1.GET("/article/list", article.List)
		v1.POST("/article/edit", article.Edit)
		v1.POST("/comment/add", comment.Add)
		v1.GET("/comment/list", comment.List)
		v1.POST("/comment/edit", comment.Edit)
	}
	return r
}
