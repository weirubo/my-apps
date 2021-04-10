package version1_0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-apps/api/cms/internal/service"
	"net/http"
)

type Tag struct {
}

type TagRes struct {
	ID        uint   `json:"id"`
	TagName   string `json:"tag_name"`
	Count     uint   `json:"count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (t Tag) Add(c *gin.Context) {
	param := &service.TagReq{}
	err := c.ShouldBind(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.AddTag(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (t Tag) List(c *gin.Context) {
	param := &service.Page{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	tags, err := svc.GetTags(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	var tag TagRes
	var tagList []TagRes
	for _, v := range tags {
		tag.ID = v.ID
		tag.TagName = v.TagName
		tag.Count = v.Count
		tag.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		tag.UpdatedAt = v.UpdatedAt.Format("2006-01-02 15:04:05")
		tagList = append(tagList, tag)
	}
	data := gin.H{
		"tags": tagList,
	}
	c.JSON(http.StatusOK, data)
}

func (t Tag) Edit(c *gin.Context) {
	param := &service.TagReq{}
	err := c.ShouldBind(&param)
	if err != nil {
		fmt.Println(err)
		return
	}
	svc := service.New()
	err = svc.UpdateTag(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
