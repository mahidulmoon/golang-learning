package handlers

import (
	"MovieRater/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		movie, err := models.GetAllMovie()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"results": movie,
			})
		}
	}
}
