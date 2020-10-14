package api

import (
	dash "github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Add Trainer Payment

func AddTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model dash.TrainerPayment
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := model.Update()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"id": model.Id})
			}
		}
	}
}

// UPDATE TrainerPayment
func UpdateTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model dash.TrainerPayment
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := model.Update()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"id": model.Id})
			}
		}
	}
}

//GET ALL TrainerPayments
func GetTrainerPayments() gin.HandlerFunc {
	return func(c *gin.Context) {
		models, err := dash.GetAllTrainerPayments()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, models)
		}

	}
}

//GET TrainerPayment BY ID
func GetTrainerPaymentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		model := dash.TrainerPayment{}
		model.Id = int64(ID)
		model, err := dash.GetTrainerPaymentById(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, model)
		}
	}
}
