package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WorkShopCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		var workshop models.WorkshopRequest

		err := c.ShouldBindJSON(&workshop)

		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error to bind" : err,
			})
		}else{
			err := workshop.AddWorkShop()
			if err != nil{
				c.JSON(http.StatusUnprocessableEntity,gin.H{
					"error to insert" : err,
				})
			}else {
				c.JSON(200,gin.H{
					"success" : "workshop addedd successfully",
				})
			}
		}
	}
}