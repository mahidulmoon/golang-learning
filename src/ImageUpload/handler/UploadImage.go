package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}

		filename := header.Filename

		fmt.Println(file, err, filename)

		out, err := os.Create("static/" + filename)
		//Note that the static/uploadfile / here is not / static/uploadfile/
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{
			"hello": "success",
		})
	}
}
