package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"net/http"
)

type User struct {
}

// 定义方法
func (u User) Add(c *gin.Context) {
	param := &service.User{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.Add(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
