package dao

import (
	"fmt"
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

func (d *Dao) CreateArticle(title, description, content, categoryName, tagName string, categoryID, tagID uint8, commentCount uint) error {
	article := model.Article{
		Model:        gorm.Model{},
		Title:        title,
		Description:  description,
		CategoryID:   categoryID,
		CategoryName: categoryName,
		CommentCount: commentCount,
	}
	articleView := model.ArticleView{
		Model:        gorm.Model{},
		Title:        title,
		Content:      content,
		CategoryID:   categoryID,
		CategoryName: categoryName,
		TagID:        tagID,
		TagName:      tagName,
		CommentCount: commentCount,
	}
	tx := d.dbEngine.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.AutoMigrate(article); err != nil {
		fmt.Println(err)
		return err
	}
	if err := tx.AutoMigrate(articleView); err != nil {
		fmt.Println(err)
		return err
	}
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&articleView).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (d *Dao) ReadArticles(pageNumber, pageSize int) ([]model.Article, error) {
	article := model.Article{}
	return article.ReadByPage(d.dbEngine, pageNumber, pageSize)
}
