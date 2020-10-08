package main

import (
	"fmt"
	db "MovieRater/databaseconnector"
	"MovieRater/handlers"
	"MovieRater/models"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	
	dbpg := db.Connect()
	models.CreateProductItemsTable(dbpg)

	v1 := viper.New()
	v1.SetConfigFile("test.yaml")
	v1.ReadInConfig()
	fmt.Println(v1.GetInt("port"))
	fmt.Println(v1.GetString("hostname"))


	router := gin.Default()

	router.Use(CORSMiddleware())

	router.POST("/userregister", handlers.UserRegister())
	router.GET("/alluser", handlers.AllUser())
	router.POST("/createmovie", handlers.MovieCreate())
	router.GET("/allmovie", handlers.GetMovie())
	router.POST("/rating", handlers.Rating())

	router.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, X-Auth-Secret, Uid, Aid, CToken")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
