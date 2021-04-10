package version1_0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"net/http"
)

type Category struct {
}

type CategoryRes struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Count        uint   `json:"count"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (cate Category) Add(c *gin.Context) {
	param := &service.CategoryReq{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.AddCategory(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (cate Category) List(c *gin.Context) {
	param := &service.Page{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	categories, err := svc.GetCategories(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	var category CategoryRes
	var categoryList []CategoryRes
	for _, v := range categories {
		category.ID = v.ID
		category.CategoryName = v.CategoryName
		category.Count = v.Count
		category.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		category.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		categoryList = append(categoryList, category)
	}

	data := gin.H{
		"categories": categoryList,
	}
	c.JSON(http.StatusOK, data)
}

func (cate Category) Edit(c *gin.Context) {
	param := &service.CategoryReq{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.UpdateCategory(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (cate Category) Remove(c *gin.Context) {
	param := &service.CategoryReq{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.DeleteCategory(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
