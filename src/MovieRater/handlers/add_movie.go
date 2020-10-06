package handlers

import (
	"MovieRater/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MovieCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.Movie

		err := c.ShouldBindJSON(&movie)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"errorjson": err,
			})
		} else {
			err := movie.AddMovie()
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"errortosave": err,
				})
			} else {
				c.JSON(200, gin.H{
					"success": "movie added successfully",
				})
			}
		}
	}
}
