package main

import (
	//"github.com/gin-gonic/gin"
	"offersapp/db"
	"offersapp/models"
)

func main() {
	//router := gin.Default()
	//usersGroup = router.Group("users")
	models.CreateLoginTable(db.DB)
}
