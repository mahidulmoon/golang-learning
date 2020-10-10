package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilesCreate() gin.HandlerFunc{
	return func(c *gin.Context){
		var file models.Files

		err := c.ShouldBindJSON(&file)

		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error to bind" : err,
			})
		}else{
			err := file.AddFiles()
			if err != nil{
				c.JSON(http.StatusUnprocessableEntity,gin.H{
					"error to insert" : err,
				})
			}else {
				c.JSON(200,gin.H{
					"success" : "files added successfully",
				})
			}
		}
	}
}