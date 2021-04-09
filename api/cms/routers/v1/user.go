package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"net/http"
)

type User struct {
}

type UserReq struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 定义方法
func (u User) Add(c *gin.Context) {
	param := &service.UserReq{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.AddUser(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (u User) GetUsers(c *gin.Context) {
	param := &service.Page{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	users, err := svc.GetUsers(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	var user UserReq
	var userList []UserReq
	for _, v := range users {
		user.ID = v.ID
		user.UserName = v.Username
		user.Email = v.Email
		user.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		user.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		userList = append(userList, user)
	}

	data := gin.H{
		"users": userList,
	}
	c.JSON(http.StatusOK, data)
}
