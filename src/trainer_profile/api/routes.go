package api

import (
	"github.com/gin-gonic/gin"

)

func RunServer(port string) {


	//Initiate router
	router := gin.Default()
	router.Use(CORSMiddleware())

	//Access groups
	admin := router.Group("/admin")
	trainer := router.Group("/trainer")
	//---------------------ADMIN
	//WORKSHOP REQUEST
	admin.GET("/workshop-requests/",GetWorkShopRequests())
	admin.GET("/worshop-request/:id",GetWorkShopResponseById())
	admin.GET("/workshop-request/status/:status",GetWorkshopRequestbyStatus())
	admin.POST("/worshop-request/:id",UpdateWorkShopRequest())
	admin.DELETE("/worshop-request/:id",DeleteWorkShopRequestById())
	admin.GET("/workshop-request/workshop/:id",GetWSRbyWID())
	//FILES
	admin.GET("/workshop-files/",GetAllFiles())
	admin.GET("/workshop-files/:id",GetFilesById())

	admin.POST("/workshop-files/",FilesCreate())
	admin.POST("/workshop-files/:id/update",UpdateFiles())
	admin.DELETE("/workshop-files/:id",DeleteFileById())

	//TRAINER PAYMENT
	admin.GET("/payments",GetTrainerPayments())
	admin.GET("/payment/:id",GetTrainerPaymentByID())
	admin.POST("/payment/:id",AdminAddTrainerPayment())
	admin.PATCH("/payment/:id",UpdateTrainerPayment())
	admin.GET("/payments/due", GetTrainersPaymentDue())//-----
	//---------------------TRAINER
	//WORKSHOP REQUEST
	trainer.GET("/worshop-requests/",GetWorkshopRequestsByTrainer())//WS by trainer ID from token
	trainer.GET("/worshop-request/:id",TrainerGetWorkShopResponseById())//TOKEN
	trainer.POST("/worshop-request/:id",WorkshopRequestCreate())//TOKEN

	//FILES
	trainer.GET("/workshop-files/ws/:id", TrainerGetFilesbyWID())//TOKEN
	trainer.POST("/workshop-files/:id", FilesCreate())//TOKEN

	//TRAINER PAYMENT
	trainer.GET("/trainer-payment/",GetTrainerPaymentByTID())//token
	trainer.POST("/trainer-payment/:id",TrainerAddTrainerPayment())//token
	trainer.POST("/trainer-payment/:id/update",UpdateTrainerPayment())//token+add comment


	router.Run(port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, X-Auth-Secret, Uid, Aid, CToken")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
