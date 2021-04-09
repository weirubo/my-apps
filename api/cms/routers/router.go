package routers

import (
	"github.com/gin-gonic/gin"
)
import v1 "my-apps/api/cms/routers/v1"

// 创建路由
func NewRouter() *gin.Engine {
	// 不使用默认中间件
	r := gin.New()
	user := v1.User{}
	category := v1.Category{}
	tag := v1.Tag{}
	article := v1.Article{}
	comment := v1.Comment{}
	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.POST("/user/add", user.Add)
		v1.POST("/category/add", category.Add)
		v1.POST("/tag/add", tag.Add)
		v1.POST("/article/add", article.Add)
		v1.POST("/comment/add", comment.Add)
	}
	return r
}
