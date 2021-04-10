package dao

import (
	"gorm.io/gorm"
	"my-apps/api/cms/internal/model"
)

func (d *Dao) CreateTag(tagName string, count uint) error {
	tag := model.Tag{
		Model:   gorm.Model{},
		TagName: tagName,
		Count:   count,
	}
	return tag.Create(d.dbEngine)
}

func (d *Dao) ReadTags(pageNumber, pageSize int) ([]model.Tag, error) {
	tag := model.Tag{}
	return tag.ReadByPage(d.dbEngine, pageNumber, pageSize)
}

func (d *Dao) UpdateTag(id uint, tagName string) error {
	tag := model.Tag{
		Model: gorm.Model{
			ID: id,
		},
		TagName: tagName,
	}
	return tag.UpdateByID(d.dbEngine)
}

func (d *Dao) DeleteTag(id uint) error {
	tag := model.Tag{
		Model: gorm.Model{
			ID: id,
		},
	}
	return tag.DeleteByID(d.dbEngine)
}
