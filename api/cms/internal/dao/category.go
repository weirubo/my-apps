package dao

import (
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

func (d *Dao) CreateCategory(categoryName string, count uint) error {
	category := model.Category{
		Model:        gorm.Model{},
		CategoryName: categoryName,
		Count:        count,
	}
	return category.Create(d.dbEngine)
}

func (d *Dao) ReadCategories(pageNumber, pageSize int) ([]model.Category, error) {
	category := model.Category{}
	return category.ReadByPage(d.dbEngine, pageNumber, pageSize)
}
