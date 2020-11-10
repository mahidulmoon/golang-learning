package api

import (
	"coversion/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLetterCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var values models.NewsLetter

		err := c.ShouldBindJSON(&values)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			err := values.Add()
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"error": "email taken or phone taken or something wrong",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "newsletter added successful",
				})
			}
		}

	}
}

func NewsLetterValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var values models.NewsLetter
		fmt.Println(values.Email)
		err := c.ShouldBindJSON(&values)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			userid, err := values.ValidateNewLetter()
			if userid == 0 {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"message": "subscription failed",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "email has been taken || subscription successful",
					"status":  "subscribed || taken",
				})
			}
		}

	}
}
