package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWorkShopRequests() gin.HandlerFunc{
	return func(c *gin.Context){
		workshop,err := models.GetWorkshopRequest()
		if err != nil{
			c.JSON(http.StatusNotFound,gin.H{
				"results" : "not found",
			})
		} else {
			c.JSON(200,gin.H{
				"result" : workshop,
			})
		}

	}
}