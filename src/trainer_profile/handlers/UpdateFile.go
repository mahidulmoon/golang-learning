package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UpdateFiles() gin.HandlerFunc{
	return func(c *gin.Context){
		var file models.Files
		err := c.ShouldBindJSON(&file)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error" : err,
			})
		}else {
			err := file.UpdateFiles()
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