package service

import "my-apps/api/cms/internal/model"

type CategoryReq struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Count        uint   `json:"count"`
}

func (s service) AddCategory(param *CategoryReq) error {
	return s.dao.CreateCategory(param.CategoryName, param.Count)
}

func (s service) GetCategories(param *Page) ([]model.Category, error) {
	return s.dao.ReadCategories(param.PageNumber, param.PageSize)
}

func (s service) UpdateCategory(param *CategoryReq) error {
	return s.dao.UpdateCategory(param.ID, param.CategoryName)
}

func (s service) DeleteCategory(param *CategoryReq) error {
	return s.dao.DeleteCategory(param.ID)
}
