package version1_0

import (
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"my-apps/api/cms/pkg"
	"net/http"
)

type User struct {
}

type UserRes struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u User) Add(c *gin.Context) {
	param := &service.UserReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.AddUser(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20101,
			"msg":  "add user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (u User) List(c *gin.Context) {
	param := &service.Page{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	users, err := svc.GetUsers(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20102,
			"msg":  "get users failed",
		})
		return
	}
	var user UserRes
	var userList []UserRes
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

func (u User) Edit(c *gin.Context) {
	param := &service.UserReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.UpdateUser(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20103,
			"msg":  "edit user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (u User) Remove(c *gin.Context) {
	param := &service.UserReq{}
	err := pkg.ShouldAndValid(c, param)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}
	svc := service.New()
	err = svc.DeleteUser(param)
	if err != nil {
		// TODO::写日志
		c.JSON(http.StatusOK, gin.H{
			"code": 20103,
			"msg":  "delete user failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
