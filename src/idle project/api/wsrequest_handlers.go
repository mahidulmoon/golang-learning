package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UpskillBD/BE-TrainerDash/models"
	"github.com/gin-gonic/gin"
)

func DeleteWorkShopRequestById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteWorkshopRequest(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "did not matched",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "deleted successfully",
			})
		}
	}
}

func GetWorkShopRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		workshop, err := models.GetWorkshopRequests()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "could not found",
			})
		} else {
			c.JSON(200, gin.H{
				"data":    workshop,
				"message": "match found",
			})
		}

	}
}

func GetWorkshopRequestbyStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")
		models, err := models.GetWorkshopRequestbyStatus(status)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "no match found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":    models,
				"message": "match found",
			})
		}

	}
}
func GetWorkShopResponseById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		workshoprequest, err := models.GetWorkshopRequestById(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not fetch data",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    workshoprequest,
			})
		}
	}
}
func GetWSRbyWID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		workshoprequest, err := models.GetWSRbyWID(ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not fetch data",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fetching data successful",
				"data":    workshoprequest,
			})
		}
	}
}

func UpdateWorkShopRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var workshop models.WorkshopRequest
		err := c.ShouldBindJSON(&workshop)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		} else {
			err := workshop.Update()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "could not update",
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{
					"message": "updated successfully",
				})
			}
		}
	}
}

//-----------------------------------------------require trainertoken
//
func TrainerGetFilesbyWID() gin.HandlerFunc {
	return func(c *gin.Context) {
		wid, _ := strconv.Atoi(c.Param("id"))
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			workshoprequest, err := models.TrainerGetWSRbyWID(wid, tid)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "workshop request is not connected by wid, contact admin.",
					"error":   err,
				})
			} else {
				c.JSON(http.StatusOK, workshoprequest.File)
			}
		}
	}
}

func GetWorkshopRequestsByTrainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			models, err := models.GetWorkshopRequestByTrainer(tid)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "could not fetch data",
				})

			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "fetching data successful",
					"data":    models,
				})

			}
		}
	}
}

func TrainerGetWorkShopResponseById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			workshoprequest, err := models.TrainerGetWorkshopRequestById(ID, tid)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "no match found",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "fetching data successful",
					"data":    workshoprequest,
				})
			}
		}

	}
}

func WorkshopRequestCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var workshop models.WorkshopRequest

		err := c.ShouldBindJSON(&workshop)
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		tid, err := strconv.Atoi(TrainerID.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "You are not authorized to access",
			})
		} else {
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": "could not bind json",
				})
			} else {
				workshop.Trainer_Id = int64(tid)
				err := workshop.Add()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusUnprocessableEntity, gin.H{
						"error": "could not insert",
					})
				} else {
					c.JSON(200, gin.H{
						"message": "workshop addedd successfully",
					})
				}
			}
		}

	}
}
