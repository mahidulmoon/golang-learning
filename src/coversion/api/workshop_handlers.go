package api

import (
	"coversion/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WorkShopCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var values models.Workshop

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
					"message": "workshop added successful",
				})
			}
		}

	}
}

func GetWorkShopRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		workshop, err := models.GetWorkshops()
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "could not fetch workshops",
			})
		} else {
			c.JSON(200, gin.H{
				"data":    workshop,
				"message": "match found",
			})
		}

	}
}

func GetWorkShopById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		workshop, err := models.GetWorkshopByID(id)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "could not fetch workshops",
			})
		} else {
			c.JSON(200, gin.H{
				"data":    workshop,
				"message": "match found",
			})
		}

	}
}
