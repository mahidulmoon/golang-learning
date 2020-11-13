package services

import (
	"github.com/UpskillBD/upskill-main/models"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "",
		password: "",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	var model models.User
	result := model.Authenticate(email, password)
	//return info.email == email && info.password == password
	if result {
		return true
	} else {
		return false
	}
}
