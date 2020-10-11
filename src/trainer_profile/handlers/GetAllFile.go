package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func AllFiles() gin.HandlerFunc{
	return func(c *gin.Context){
		
		files,err := models.GetAllFiles()
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error": err,
			})
		}else{
			c.JSON(200,gin.H{
				"results": files,
			})
		}
	}
}