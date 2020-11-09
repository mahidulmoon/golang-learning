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
				c.JSON(500, gin.H{
					"error": "email taken or phone taken or something wrong",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "registration successful",
				})
			}
		}

	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header.Get("X-Auth-Secret")
		// TrainerID, err := Instance.GetUserFromToken(token)
		ID := 2
		user, err := models.GetUserById(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"error": "could not fetch user request || invalid token",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":    user,
				"message": "resull found",
			})
		}
	}
}
