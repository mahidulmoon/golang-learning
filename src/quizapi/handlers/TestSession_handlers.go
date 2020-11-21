package handlers

import (
	"domaniGoupdate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		var model models.TestSession
		err := c.ShouldBindJSON(&model)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unmarshal failed",
				"error":   err,
			})
		} else {
			id, err := model.NewSession()
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
