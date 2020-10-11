package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)


func DeleteFileById() gin.HandlerFunc{
	return func(c *gin.Context){
		ID,_ := strconv.Atoi(c.Param("id"))
		
		err := models.DeleteFiles(ID)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error": err,
			})
		}else{
			c.JSON(200,gin.H{
				"success":"deleted successfully",
			})
		}
	}
}