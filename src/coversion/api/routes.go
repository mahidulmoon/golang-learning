package api

import (
	"github.com/gin-gonic/gin"
)

func Runserver(port string) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	//usertable
	router.POST("/api/register", UserCreate())
	router.GET("/api/user", GetUserById())
	router.POST("/api/validate", Validation())

	//workshop
	router.POST("/api/addworkshop", WorkShopCreate())
	router.GET("api/workshops", GetWorkShopRequests())
	router.GET("api/workshops/:id", GetWorkShopById())

	//newsletter
	router.POST("/api/addnewsletter", NewLetterCreate())
	router.POST("/api/newsletter/validate", NewsLetterValidation())

	router.Run(port)
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
