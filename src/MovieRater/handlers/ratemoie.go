package handlers

import (
	"MovieRater/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Rating() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rating models.Rating

		err := c.ShouldBindJSON(&rating)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err,
			})
		} else {
			err := rating.RateMovie()

			if err != nil {
				c.JSON(http.StatusNotImplemented, gin.H{
					"error": err,
				})
			} else {
				c.JSON(200, gin.H{
					"success": "rating added successfully",
				})
			}
		}
	}
}
