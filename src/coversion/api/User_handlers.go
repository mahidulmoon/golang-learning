package api

import (
	"github.com/gin-gonic/gin"
)

func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "success",
		})
	}
}
