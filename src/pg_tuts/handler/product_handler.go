package handler

import(
	"github.com/gin-gonic/gin"
)

func GetAllProduct() gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(200, gin.H{
			"hello": "found me",
		})
	}
}