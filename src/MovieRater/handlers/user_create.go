package handlers

import(
	"github.com/gin-gonic/gin"
	"MovieRater/models"
	"net/http"
	// "strconv"
)


func UserRegister() gin.HandlerFunc{
	return func(c *gin.Context){
		var model models.User
		err := c.ShouldBindJSON(&model)

		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err,
			})
		}else{
			err := model.CreateUser()
			
			if err != nil{
				c.JSON(200,gin.H{
					"error": err,
				})
			} else {
				c.JSON(200,gin.H{
					"success":"user added successfully",
				})
			}
		}

	}
}