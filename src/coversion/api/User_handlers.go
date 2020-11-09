package api

import (
	"coversion/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var values models.User

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
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not insert",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "user addedd successfully",
				})
			}
		}

	}
}
