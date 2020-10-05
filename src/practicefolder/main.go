package main

import(
	"fmt"
	db "practicefolder/databaseconnector"
	"github.com/gin-gonic/gin"
	"practicefolder/handlers"
	"practicefolder/models"
)

func main(){
	fmt.Println("hello")
	dbpg := db.Connect()
	models.CreateProductItemsTable(dbpg)
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/products",handlers.GetMeeting())
	router.POST("/products",handlers.Add())
	router.GET("/products/:id",handlers.GetProductBytid())
	

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