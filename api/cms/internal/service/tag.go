package service

type Tag struct {
	TagName string `json:"tag_name"`
	Count   uint   `json:"count"`
}

func (s service) AddTag(param *Tag) error {
	return s.dao.CreateTag(param.TagName, param.Count)
}
