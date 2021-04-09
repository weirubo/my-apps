package service

type Category struct {
	CategoryName string `json:"category_name"`
	Count        uint   `json:"count"`
}

func (s service) AddCategory(param *Category) error {
	return s.dao.CreateCategory(param.CategoryName, param.Count)
}
