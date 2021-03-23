package service

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (s service) Add(param *User) error {
	return s.dao.Create(param.UserName, param.Email, param.PassWord)
}
