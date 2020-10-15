package api

import (
	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Add Trainer Payment

func AdminAddTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
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


func TrainerAddTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)
		//model.Trainer_Id=token
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
		var model models.TrainerPayment
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
func TrainerUpdateTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := model.TrainerUpdate()
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
		models, err := models.GetAllTrainerPayments()
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
		model := models.TrainerPayment{}
		model.Id = int64(ID)
		model, err := models.GetTrainerPaymentById(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, model)
		}
	}
}

func GetTrainersPaymentDue()gin.HandlerFunc{
	return func(c *gin.Context) {
		models, err := models.GetTrainersPaymentDue()
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, models)
		}

	}
}


//--------------------------------token
func GetTrainerPaymentByTID()gin.HandlerFunc{
	return func(c *gin.Context) {
		tid:=123 //token
		file, err := models.GetTrainerPaymentByTID(int64(tid))
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": file,
			})
		}


	}
}