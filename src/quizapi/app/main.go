package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"domaniGoupdate/handlers"
	"domaniGoupdate/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/domains", handlers.GetDomains())

	r.GET("/subdomains/:id", handlers.GetSubdomainByID())

	r.GET("/topics/:id", handlers.GetTopicsByID())

	r.GET("/tests/:id", handlers.GetTestByID())

	r.POST("/submit-result", handlers.SubmitResult())

	r.GET("/user-stats", handlers.GetStats())

	r.POST("/domains", handlers.AddDomain())

	r.POST("/subdomains/", handlers.AddSubdomain())

	r.POST("/topics/", handlers.AddTopic())

	r.POST("/test", handlers.AddTest())

	r.POST("/session", handlers.AddSession())

	//---Useless, Consider for deletion
	/*r.GET("/root", func(c *gin.Context) {
		items, err := models.GetRoot()
		if err == nil {
			c.JSON(http.StatusOK, items)
			return
		}
		c.String(http.StatusNotFound, "could not document")
	})
	r.PUT("/root", func(c *gin.Context) {
		request := struct {
			Data []models.Domain `json:"data"`
		}{}
		err := c.ShouldBindJSON(&request)
		for key := range request.Data {
			err = request.Data[key].Update()
		}
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "updated",
			})
		}
		c.String(http.StatusNotFound, "could not document")
	})
	*/
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "invalid file")
			return
		}
		OriginalFilename := strings.Split(filepath.Base(file.Filename), ".")
		filename := models.NewUUID() + "." + OriginalFilename[len(OriginalFilename)-1]
		fileSrc := "./uploads/" + filename
		if err := c.SaveUploadedFile(file, fileSrc); err != nil {
			c.JSON(http.StatusBadRequest, "could not upload file")
			return
		}
		c.JSON(200, gin.H{
			"url": fileSrc,
		})
	})

	r.Run(":9001")
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
