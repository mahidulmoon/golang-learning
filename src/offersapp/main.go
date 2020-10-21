package main

import (
	"github.com/gin-gonic/gin"
	//"offersapp/db"
	//"offersapp/models"
	"offersapp/routes"
)

func main() {
	router := gin.Default()
	//models.CreateLoginTable(db.DB)

	usersGroup := router.Group("/users")

	usersGroup.POST("/register", routes.UsersRegister)

	router.Run()
}
