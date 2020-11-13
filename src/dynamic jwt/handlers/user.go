package handlers

import (
	"fmt"
	"net/http"

	"github.com/UpskillBD/upskill-main/models"

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

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var values models.User
		fmt.Println(values.Email)
		err := c.ShouldBindJSON(&values)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			userid, err := values.Validate()
			if userid == 0 {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"message": "validation failed",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "email has been taken || phone number has been taken",
				})
			}
		}

	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.User
		token := c.Request.Header.Get("Bearer")
		TrainerID, err := model.GetUserfromToken(token)
		//ID := 2 //get id from token
		user, err := models.GetUserById(int(TrainerID))
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
