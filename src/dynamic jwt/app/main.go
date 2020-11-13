package main

import (
	"github.com/UpskillBD/upskill-main/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(handlers.CORSMiddleware())

	api := router.Group("/api")

	//login
	api.POST("/login", handlers.UserLogin())

	//workshops
	api.GET("/workshops", handlers.GetWorkshops())
	api.GET("/workshops/:id", handlers.GetWorkShopById())
	api.POST("/workshops", handlers.CreateWorkshop())
	api.PUT("/workshops", handlers.UpdateWorkshop())
	api.POST("/upload", handlers.FileUpload())
	api.GET("/workshops/:id/join", handlers.JoinWS())                  //token
	api.GET("/workshops/:id/cancel", handlers.CancelWS())              //token
	api.GET("/workshops/:id/status", handlers.GetWorkShopStatusById()) //token

	//user
	api.POST("/register", handlers.UserCreate())
	api.POST("/validate", handlers.Validation())
	api.GET("/user", handlers.GetUserById()) //token

	//newsletter
	api.POST("/newsletter/validate", handlers.NewsLetterValidation())

	router.Run(":8083")
}
