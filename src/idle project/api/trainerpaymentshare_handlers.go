package api

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
)

func TrainerPaymentShareCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var share models.TrainerPaymentShare
		err := c.ShouldBindJSON(&share)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			err := share.Add()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not insert share",
				})
			} else {

				c.JSON(http.StatusCreated, gin.H{

					"message": "TrainerPaymentShare added successfully",
				})
			}
		}
	}
}

func GetAllTrainerPaymentShare() gin.HandlerFunc {
	return func(c *gin.Context) {
		shares, err := models.GetAllShare()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "could not fetch data",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    shares,
			})

		}

	}
}

func GetShareByWID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("wid"))

		share, err := models.GetShareWid(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not fetch data",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    share,
			})
		}

	}
}

func UpdateTrainerPaymentShare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var share models.TrainerPaymentShare
		err := c.ShouldBindJSON(&share)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			err := share.UpdateShare()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": "could not update share",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"message": "updated successfully",
				})
			}
		}
	}
}

func DeleteShareByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))

		err := models.DeleteShareid(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not delete",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete successfully",
			})
		}

	}
}
