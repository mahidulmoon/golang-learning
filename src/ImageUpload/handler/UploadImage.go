package handler

import (
	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "success",
		})
	}
}
