package handlers

import(
	"github.com/UpskillBD/trainer-dashboard-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

func GetWorkShopResponseById() gin.HandlerFunc{
	return func(c *gin.Context){
		ID,_ := strconv.Atoi(c.Param("id"))
		workshoprequest,err := models.GetWorkshopRequestById(ID)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error": err,
			})
		}else{
			c.JSON(http.StatusOK,workshoprequest)
		}
	}
}