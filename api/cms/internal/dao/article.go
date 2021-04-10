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
	// TODO::更新分类表的文章总数和标签表的文章总数
	return tx.Commit().Error
}

func (d *Dao) ReadArticles(pageNumber, pageSize int) ([]model.Article, error) {
	article := model.Article{}
	return article.ReadByPage(d.dbEngine, pageNumber, pageSize)
}

func (d *Dao) UpdateArticle(id, commentCount uint, title, description, content, categoryName, tagName string, categoryID, tagID uint8) error {
	article := model.Article{
		Model: gorm.Model{
			ID: id,
		},
		Title:        title,
		Description:  description,
		CategoryID:   categoryID,
		CategoryName: categoryName,
		CommentCount: commentCount,
	}

	articleView := model.ArticleView{
		Model: gorm.Model{
			ID: id,
		},
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
	if err := tx.Updates(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Updates(&articleView).Error; err != nil {
		tx.Rollback()
		return err
	}
	// TODO::更新分类表的文章总数和标签表的文章总数
	return tx.Commit().Error
}

func (d *Dao) DeleteArticle(id uint) error {
	article := model.Article{
		Model: gorm.Model{
			ID: id,
		},
	}
	articleView := model.ArticleView{
		Model: gorm.Model{
			ID: id,
		},
	}
	tx := d.dbEngine.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Delete(&article, &article.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&articleView, &articleView.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	// TODO::更新分类表的文章总数和标签表的文章总数
	return tx.Commit().Error
}

func (d *Dao) ReadArticle(id uint) (model.ArticleView, error) {
	articleView := model.ArticleView{
		Model: gorm.Model{
			ID: id,
		},
	}
	return articleView.Read(d.dbEngine)
}
