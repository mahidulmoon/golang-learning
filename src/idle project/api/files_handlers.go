package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
)

func FilesCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file models.Files

		err := c.ShouldBindJSON(&file)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   "could not bind",
				"message": "could not bind",
			})
		} else {
			resp, err := file.Add()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error":   "could not create",
					"message": "could not create",
				})
			} else {

				c.JSON(http.StatusCreated, gin.H{
					"data":    resp,
					"message": "files added successfully",
				})
			}
		}
	}
}

func TrainerFilesCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file models.Files

		err := c.ShouldBindJSON(&file)
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			file.Trainer_Id = int64(tid)

			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not bind",
				})
			} else {
				resp, err := file.Add()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"error":   "could not insert",
						"message": "could not insert",
					})
				} else {

					c.JSON(http.StatusCreated, gin.H{
						"data":    resp,
						"message": "files added successfully",
					})
				}
			}
		}
	}
}
func DeleteFileById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteFile(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   "couldn't delete",
				"message": "couldn't delete",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "deleted successfully",
			})
		}
	}
}

func GetAllFiles() gin.HandlerFunc {
	return func(c *gin.Context) {

		files, err := models.GetAllFiles()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not get all files",
			})
		} else {
			c.JSON(200, gin.H{
				"data":    files,
				"message": "files loaded successful",
			})
		}
	}
}

func GetFilesById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		file, err := models.GetFilesByID(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":    file,
				"message": "resull found",
			})
		}
	}
}

func UpdateFiles() gin.HandlerFunc {
	return func(c *gin.Context) {
		var file models.Files
		err := c.ShouldBindJSON(&file)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind",
			})
		} else {
			resp, err := file.Update()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"data":  resp,
					"error": "could not update",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"message": "updated successfully"})
			}
		}
	}
}
