package service

import "my-apps/api/cms/internal/model"

type TagReq struct {
	TagName string `json:"tag_name"`
	Count   uint   `json:"count"`
}

func (s service) AddTag(param *TagReq) error {
	return s.dao.CreateTag(param.TagName, param.Count)
}

func (s service) GetTags(param *Page) ([]model.Tag, error) {
	return s.dao.ReadTags(param.PageNumber, param.PageSize)
}
