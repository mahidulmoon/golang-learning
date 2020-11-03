package api
import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//List of valid extensions
var extensions []string = []string{".pdf", ".doc", ".docx", ".odt", ".txt", ".ppt", ".pptx", ".odp", ".xls", ".xlsx", ".ods", ".jpg", ".jpeg", ".png", ".bmp", ".gif"}

// IsValidExtension checks whether if the extension is valid for upload
func IsValidExtension(extension string) bool {
	for _, ext := range extensions {
		if ext == extension {
			return true
		}
	}
	return false
}

// GetFileIcon returns an icon for the file
func GetFileIcon(extension string) (string, error) {
	for _, ext := range []string{".pdf", ".doc", ".docx", ".odt", ".txt"} {
		if ext == extension {
			return "document.png", nil
		}
	}
	for _, ext := range []string{".ppt", ".pptx", ".odp"} {
		if ext == extension {
			return "presentation.png", nil
		}
	}
	for _, ext := range []string{".xls", ".xlsx", ".ods"} {
		if ext == extension {
			return "spreadsheet.png", nil
		}
	}
	for _, ext := range []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"} {
		if ext == extension {
			return "image.png", nil
		}
	}
	return "file.png", errors.New("invalid extension")
}

func UploadHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		// Exits if uploaded file is not a valid extension
		if !IsValidExtension(filepath.Ext(file.Filename)) {
			fmt.Println("INVALID EXT")
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: Invalid file type"))
			return
		}
		img, err := GetFileIcon(filepath.Ext(file.Filename))

		OriginalFilename := strings.Split(filepath.Base(file.Filename), ".")
		filename := fmt.Sprint(time.Now().Unix()) + "." + OriginalFilename[len(OriginalFilename)-1]
		//fmt.Println(filename)
		fileSrc := "../uploads/" + filename
		imgSrc := "../icons/" + img
		// FILE SAVED
		if err := c.SaveUploadedFile(file, fileSrc); err != nil {
			//fmt.Println("Something went wrong!!")
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.JSON(200, gin.H{
			"url":  fileSrc,
			"icon": imgSrc,
		})
	}
}

