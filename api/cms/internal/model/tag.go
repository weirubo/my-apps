package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(20) not null;commit:标签名称"`
	Count   uint   `gorm:"type:int(10) not null;default:0;commit:文章总数"`
}

func (t Tag) Create(db *gorm.DB) error {
	err := db.AutoMigrate(t)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return db.Create(&t).Error
}

func (t Tag) ReadByPage(db *gorm.DB, pageNumber, pageSize int) ([]Tag, error) {
	var tags []Tag
	pageOffset := (pageNumber - 1) * pageSize
	tx := db.Select("id", "tag_name", "count", "created_at", "updated_at").Limit(pageSize).Offset(pageOffset).Find(&tags)
	if err := tx.Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tags, nil
}

func (t Tag) UpdateByID(db *gorm.DB) error {
	tx := db.Where("id", t.ID).Updates(&t)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (t Tag) DeleteByID(db *gorm.DB) error {
	tx := db.Delete(&t, t.ID)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
