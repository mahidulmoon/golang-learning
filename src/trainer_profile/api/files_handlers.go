package api

import (
	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FilesCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file models.Files

		err := c.ShouldBindJSON(&file)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error to bind": err,
			})
		} else {
			resp,err := file.Add()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error to insert": err,
				})
			} else {

				c.JSON(http.StatusCreated, gin.H{
					"file": resp,
					"success": "files added successfully",
				})
			}
		}
	}
}
func DeleteFileById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteFile(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"success": "deleted successfully",
			})
		}
	}
}

func GetAllFiles() gin.HandlerFunc {
	return func(c *gin.Context) {

		files, err := models.GetAllFiles()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"results": files,
			})
		}
	}
}

func GetFilesById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		file, err := models.GetFilesByID(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": file,
			})
		}
	}
}

func UpdateFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file models.Files
		err := c.ShouldBindJSON(&file)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := file.Update()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": err,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"status": "updated successfully"})
			}
		}
	}
}

