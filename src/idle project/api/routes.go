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
	//admin.GET("/workshoprequest/status/:status",GetWorkshopRequestbyStatus())
	admin.POST("/worshop-request/:id",UpdateWorkShopRequest())
	//admin.POST("/workshop-request/status/id",ChangeStatus())
	admin.DELETE("/worshop-request/:id",DeleteWorkShopRequestById())
	//FILES
	//admin.GET("/workshop-files/",GetAllFiles())
	admin.GET("/workshop-files/:id",GetFilesById())
	//admin.GET("/workshop-files/ws/:id",GetFilesByWID())
	admin.POST("/workshop-files/",FilesCreate())
	admin.POST("/workshop-files/:id/update",UpdateFiles())
	admin.DELETE("/workshop-files/:id",DeleteFileById())

	//TRAINER PAYMENT
	admin.GET("/trainer-payments/",GetTrainerPayments())
	//admin.GET("/trainer-payment/:id",GetTrainerPaymentByTID)
	//admin.GET("/trainer-payment/due", GetTrainersPaymentDue())
	admin.POST("/trainer-payment/:id",AddTrainerPayment())
	admin.POST("/trainer-payment/:id/update",UpdateTrainerPayment())

	//---------------------TRAINER
	//WORKSHOP REQUEST
	//trainer.GET("/worshop-requests/",GetWorkshopeRequestByTrainer())//WS by trainer ID from token
	trainer.GET("/worshop-request/:id",GetWorkShopResponseById())//TOKEN
	trainer.POST("/worshop-request/:id",WorkShopCreate())//TOKEN

	//FILES
	trainer.GET("/workshop-files/:id", GetFilesById())//TOKEN
	//trainer.GET("/workshop-files/ws/:id", GetFilesByWID())//TOKEN
	trainer.POST("/workshop-files/:id", FilesCreate())//TOKEN

	//TRAINER PAYMENT
	//trainer.GET("/trainer-payment/:id",GetTrainerPaymentByTID)//token
	trainer.POST("/trainer-payment/:id",AddTrainerPayment())//token
	trainer.POST("/trainer-payment/:id/update",UpdateTrainerPayment())//token+add comment



	////workshoprequest_api
	//router.POST("/addworkshop", handlers.WorkShopCreate())
	//router.GET("/allworkshoprequest", handlers.GetWorkShopRequests())
	//router.GET("/workshoprequest/:id", handlers.GetWorkShopResponseById())
	//router.DELETE("/deleteworkshoprequest/:id", handlers.DeleteWorkShopResponseById())
	//router.PATCH("/updateworkshoprequest/:id", handlers.UpdateWorkShopResponse())
	//
	////files_api
	//router.POST("/addfiles", handlers.FilesCreate())
	//router.DELETE("/deletefiles/:id", handlers.DeleteFileById())
	//router.GET("/allfiles", handlers.AllFiles())
	//router.GET("/files/:id", handlers.GetFilesById())
	//router.PATCH("/updatefiles/:id", handlers.UpdateFiles())

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
