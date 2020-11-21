package handlers

import (
	"domaniGoupdate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTestByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type param struct {
			ID string `uri:"id" binding:"required"`
		}
		var uriParam param
		err := c.ShouldBindUri(&uriParam)
		items, err := models.GetTestsByTopicID(uriParam.ID)
		if err == nil {
			c.JSON(http.StatusOK, items)
			return
		}
		c.String(http.StatusNotFound, "could not fetch topics")
	}

}

func AddTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TestNest
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unmarshal failed",
				"error":   err,
			})
		} else {
			id, err := model.NewTest()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
					"id":    id,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"group": id})
			}
		}
	}

}
