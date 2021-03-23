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
	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.POST("/user/add", user.Add)
	}
	return r
}
