package main

import (
	"pg_test/db"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	db := db.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	   })
    //r.GET("/books", controllers.FindBooks)
	
	router.Run()
}