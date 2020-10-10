package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UpdateWorkShopResponse() gin.HandlerFunc{
	return func(c *gin.Context){
		var workshop models.WorkshopRequest
		err := c.ShouldBindJSON(&workshop)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error" : err,
			})
		}else {
			err := workshop.UpdateWorkShop()
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{
					"err":err,
				})
			}else{
				c.JSON(http.StatusCreated,gin.H{"status":"updated successfully"})
			}
		}
	}
}