package handlers

import(
	"github.com/gin-gonic/gin"
	"practicefolder/models"
)

func GetMeeting() gin.HandlerFunc{
	return func(c *gin.Context){
		model, err := models.GetProducts()
		if err != nil{
			c.JSON(200, gin.H{
				"resuls": "no result found",
			})
		}else{
			c.JSON(200, gin.H{
				"resuls": model,
			})
		}
	}
}