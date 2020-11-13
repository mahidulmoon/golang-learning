package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

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

func FileUpload() gin.HandlerFunc  {
	return func (c *gin.Context)  {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		OriginalFilename := strings.Split(filepath.Base(file.Filename), ".")
		filename := fmt.Sprint(time.Now().Unix())+"." +OriginalFilename[len(OriginalFilename)-1]
		//fmt.Println(filename)
		fileSrc := "../uploads/"+filename
		if runtime.GOOS == "windows" {
			fileSrc = "../uploads/"+filename
		}
		// FILE SAVED
		if err := c.SaveUploadedFile(file, fileSrc);
			err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.JSON(200, gin.H{
			"url":"/public/workshop_cover/"+filename,
		})
	}
}
