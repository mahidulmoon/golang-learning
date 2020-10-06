package handlers

import(
	"github.com/gin-gonic/gin"
	"MovieRater/models"
	"net/http"
	// "strconv"
)


func AllUser() gin.HandlerFunc{
	return func(c *gin.Context){
		model,err := models.GetAllUser()
		if err != nil{
			c.JSON(http.StatusNotFound,gin.H{
				"results" : "not found",
			})
		} else {
			c.JSON(200,gin.H{
				"result" : model,
			})
		}
	}
}