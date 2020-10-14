package api

import (
	dash "github.com/UpskillBD/BE-TrainerDash/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func WorkShopCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var workshop dash.WorkshopRequest

		err := c.ShouldBindJSON(&workshop)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error to bind": err,
			})
		} else {
			err := workshop.Add()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error to insert": err,
				})
			} else {
				c.JSON(200, gin.H{
					"success": "workshop addedd successfully",
				})
			}
		}
	}
}

func DeleteWorkShopRequestById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := dash.DeleteWorkshopRequest(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "deleted successfully",
			})
		}
	}
}

func GetWorkShopRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		workshop, err := dash.GetWorkshopRequests()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"results": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"result": workshop,
			})
		}

	}
}

func GetWorkShopResponseById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		workshoprequest, err := dash.GetWorkshopRequestById(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, workshoprequest)
		}
	}
}

func UpdateWorkShopRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var workshop dash.WorkshopRequest
		err := c.ShouldBindJSON(&workshop)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := workshop.Update()
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
