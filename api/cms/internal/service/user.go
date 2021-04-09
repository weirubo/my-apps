package service

type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (s service) AddUser(param *User) error {
	return s.dao.CreateUser(param.UserName, param.Email, param.PassWord)
}
