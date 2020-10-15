package api

import (
	"github.com/UpskillBD/BE-TrainerDash/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func DeleteWorkShopRequestById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := models.DeleteWorkshopRequest(ID)
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
		workshop, err := models.GetWorkshopRequests()
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

func GetWorkshopRequestbyStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("id")
		models, err := models.GetWorkshopRequestbyStatus(status)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, models)
		}

	}
}
func GetWorkShopResponseById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		workshoprequest, err := models.GetWorkshopRequestById(ID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, workshoprequest)
		}
	}
}
func GetWSRbyWID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		workshoprequest, err := models.GetWSRbyWID(ID)
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
		var workshop models.WorkshopRequest
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


//-----------------------------------------------require trainertoken
//
func TrainerGetFilesbyWID() gin.HandlerFunc {
	return func(c *gin.Context) {
		wid, _ := strconv.Atoi(c.Param("wid"))
		var tid int //from token

			workshoprequest, err := models.TrainerGetWSRbyWID(wid,tid)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, workshoprequest.File)
			}




	}
}

func GetWorkshopRequestsByTrainer() gin.HandlerFunc{
	return func(c *gin.Context) {
		var tid int64 //from token
		models,err:=models.GetWorkshopeRequestByTrainer(tid)
		if err!=nil{
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})

		}else{
				c.JSON(http.StatusOK, models)


		}


	}
}

func TrainerGetWorkShopResponseById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		tid:=12 //token
			workshoprequest, err := models.TrainerGetWorkshopRequestById(ID,tid)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, workshoprequest)
			}

	}
}


func WorkshopRequestCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var workshop models.WorkshopRequest

		err := c.ShouldBindJSON(&workshop)
		//workshop.Trainer_Id=TID from token

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