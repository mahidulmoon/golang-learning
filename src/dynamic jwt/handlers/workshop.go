package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/UpskillBD/upskill-main/models"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func GetWorkshops() func(c *gin.Context) {
	return func(c *gin.Context) {
		items, err := models.GetWorkshops()
		if err == nil {
			c.JSON(200, gin.H{
				"data": items,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong!",
				"error":   err.Error(),
			})
		}
	}
}
func timeConv(str string) time.Time {
	layout := "2006-01-02T15:04-07:00"
	//str := "2020-11-24T02:15+06:00"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	return t
}
func CreateWorkshop() func(c *gin.Context) {
	return func(c *gin.Context) {
		var model models.Workshop
		var data map[string]interface{}
		err := c.ShouldBindJSON(&data)
		fmt.Println(data)
		mapstructure.Decode(data, &model)
		model.CourseOutline = data["course_outline"].(string)
		model.WorkshopTS = data["workshop_ts"].(string)
		model.CreatedBy = 1
		model.ID = 0
		model.CreatedAt = time.Now()
		model.UpdatedAt = time.Now()
		err = model.Add()
		fmt.Println(model)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"message": "created successfully!",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong!",
				"error":   err.Error(),
			})
		}
	}
}

func UpdateWorkshop() func(c *gin.Context) {
	return func(c *gin.Context) {
		var model models.Workshop
		err := c.ShouldBindJSON(&model)
		model.CreatedAt = time.Now()
		err = model.Update()
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"message": "updated successfully!",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong!",
				"error":   err.Error(),
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
			if workshop == nil {
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
}

func JoinWS() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.User
		token := c.Request.Header.Get("Bearer")
		_, err := model.GetUserfromToken(token)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "token is not valid",
			})
		} else {
			id, _ := strconv.Atoi(c.Param("id"))
			err := models.WSJoinReq(id)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"message": "could not accept the join request",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "joined successfully",
				})
			}
		}

	}
}

func CancelWS() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.User
		token := c.Request.Header.Get("Bearer")
		_, err := model.GetUserfromToken(token)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "token is not valid",
			})
		} else {
			id, _ := strconv.Atoi(c.Param("id"))
			err := models.WSCancelReq(id)
			if err != nil {
				fmt.Println(err)
				c.JSON(500, gin.H{
					"message": "could not accept the cancel request",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "cancelled successfully",
				})
			}
		}

	}
}

func GetWorkShopStatusById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.User
		token := c.Request.Header.Get("Bearer")
		_, err := model.GetUserfromToken(token)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "token is not valid",
			})
		} else {
			id, _ := strconv.Atoi(c.Param("id"))
			result, err := models.GetWorkshopStatusByID(id)
			if err != nil {
				//fmt.Println(result)
				c.JSON(500, gin.H{
					"message": "could not fetch request",
				})
			} else {
				c.JSON(500, gin.H{
					"status": result,
				})
			}
		}

	}
}
