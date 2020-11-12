package handlers

import (
	"net/http"

	"github.com/UpskillBD/upskill-main/controller"
	"github.com/UpskillBD/upskill-main/services"
	"github.com/gin-gonic/gin"
)

func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginService services.LoginService = services.StaticLoginService()
		var jwtService services.JWTService = services.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}

	}
}
