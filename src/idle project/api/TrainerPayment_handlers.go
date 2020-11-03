package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
)

// Add Trainer Payment

func AdminAddTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind",
			})
		} else {
			err := model.Update()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "could not update",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"id":      model.Id,
					"message": "update successful",
				})
			}
		}
	}
}

func TrainerAddTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)

		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"error":   "could not add payment",
				"message": "You are not authorized to access",
			})
		} else {
			model.Trainer_Id = int64(tid)

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not found user with this id",
				})
			} else {
				err := model.Add()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "could not add payment",
					})
				} else {
					c.JSON(http.StatusCreated, gin.H{
						"id":      model.Id,
						"message": "payment add successfull",
					})
				}
			}
		}

	}
}

// UPDATE TrainerPayment
func AdminUpdateTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind input",
			})
		} else {
			err := model.Update()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "could not update",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"id":      model.Id,
					"message": "payement update successfully",
				})
			}
		}
	}
}

func UpdateTrainerPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TrainerPayment
		err := c.ShouldBindJSON(&model)

		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			model.Trainer_Id = int64(tid)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "no such trainer matched",
				})
			} else {
				err := model.Update()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "could not update trainerpayment",
					})
				} else {
					c.JSON(http.StatusCreated, gin.H{
						"id":      model.Id,
						"message": "update trainerpayment successfully",
					})
				}
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
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind",
			})
		} else {
			err := model.TrainerUpdate()
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "could not update trainerpayment",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"id":      model.Id,
					"message": "trainerpayment update successfully",
				})
			}
		}
	}
}

//GET ALL TrainerPayments
func GetTrainerPayments() gin.HandlerFunc {
	return func(c *gin.Context) {
		models, err := models.GetAllTrainerPayments()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   "could not fetch data",
				"message": "error to fetch data",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    models,
			})
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
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "error to fetch data",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    model,
			})
		}
	}
}

func GetTrainersPaymentDue() gin.HandlerFunc {
	return func(c *gin.Context) {
		models, err := models.GetTrainersPaymentDue()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not fetch due payment",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    models,
			})
		}

	}
}

//--------------------------------token
func GetTrainerPaymentByTID() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not autorized to access",
			})
		} else {
			model, err := models.GetTrainerPaymentByTID(tid)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not fetch data",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"data":    model,
					"message": "fetching data successful",
				})
			}
		}

	}
}
func GetTrainerDuePaymentByTID() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not autorized to access",
			})
		} else {
			model, err := models.GetTrainerPaymentByTID(tid)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not fetch data",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"data":    model.Due,
					"message": "fetching data successful",
				})
			}
		}

	}
}

//All registrants for trainer ,total earned raw , total earned calculated
