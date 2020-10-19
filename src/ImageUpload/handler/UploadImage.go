package handler

import (
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
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

		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		m := resize.Resize(795, 636, img, resize.Lanczos3)

		out, err := os.Create("static/WSR" + filename)
		//Note that the static/uploadfile / here is not / static/uploadfile/
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		// _, err = io.Copy(out, file)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		jpeg.Encode(out, m, nil)

		c.JSON(200, gin.H{
			"hello": "success",
		})
	}
}
