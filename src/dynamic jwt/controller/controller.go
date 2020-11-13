package controller

import (
	"github.com/UpskillBD/upskill-main/dtos"
	"github.com/UpskillBD/upskill-main/models"
	"github.com/UpskillBD/upskill-main/services"
	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func LoginHandler(loginService services.LoginService,
	jWtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dtos.LoginCredentials
	var model models.User
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		token := controller.jWtService.GenerateToken(credential.Email, true)
		err := model.SetUserToken(token, credential.Email)
		if err == nil {
			return token
		}

	}
	return ""
}
