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
