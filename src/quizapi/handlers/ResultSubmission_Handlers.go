package handlers

import (
	"domaniGoupdate/OldBackend"
	"domaniGoupdate/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitResult() gin.HandlerFunc {
	return func(c *gin.Context) {
		var submission models.SubmissionModel
		c.ShouldBindJSON(&submission)
		submission.Token = c.GetHeader("Authorization")
		if len(submission.Token) > 0 {
			c.Status(http.StatusCreated)
			user, err := OldBackend.Instance.GetUserFromToken(submission.Token)
			if err == nil {
				fmt.Println(fmt.Sprintf("id is %d", user.Id))
			} else {
				fmt.Println(err.Error())
			}
			err = models.SubmitResult(user.Id, submission)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "could not insert into database",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "success",
				})
			}

		} else {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid data"})
			return
		}
	}
}

func GetStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) > 0 {
			user, err := OldBackend.Instance.GetUserFromToken(token)
			average, err := models.GetStats(user.Id)
			if err != nil {
				fmt.Println(fmt.Sprintf("id is %d", user.Id))
			} else {
				c.JSON(http.StatusOK, gin.H{"average": average})
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "no results found"})
			return
		}
	}
}
