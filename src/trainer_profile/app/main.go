package main

import (
	//"github.com/spf13/cobra"
	//"fmt"
	db "github.com/UpskillBD/trainer-dashboard-backend/db"
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/UpskillBD/trainer-dashboard-backend/handlers"
	
)

func main() {
	
	dbpg := db.Connect()
	models.CreateWorkShopRequestTable(dbpg)
	router := gin.Default()

	router.Use(CORSMiddleware())
	//workshoprequest_api
	router.POST("/addworkshop", handlers.WorkShopCreate())
	router.GET("/allworkshoprequest", handlers.GetWorkShopRequests())
	router.GET("/workshoprequest/:id", handlers.GetWorkShopResponseById())
	router.DELETE("/deleteworkshoprequest/:id", handlers.DeleteWorkShopResponseById())
	router.PATCH("/updateworkshoprequest/:id", handlers.UpdateWorkShopResponse())
	
	//files_api
	router.POST("/addfiles", handlers.FilesCreate())
	router.DELETE("/deletefiles/:id", handlers.DeleteFileById())
	router.GET("/allfiles", handlers.AllFiles())
	router.GET("/files/:id", handlers.GetFilesById())
	router.PATCH("/updatefiles/:id", handlers.UpdateFiles())

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