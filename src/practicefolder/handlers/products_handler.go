package handlers

import(
	"github.com/gin-gonic/gin"
	"practicefolder/models"
	"net/http"
)

func GetMeeting() gin.HandlerFunc{
	return func(c *gin.Context){
		model, err := models.GetProducts()
		if err != nil{
			c.JSON(200, gin.H{
				"resuls": "no result found",
			})
		}else{
			c.JSON(200, gin.H{
				"resuls": model,
			})
		}
	}
}


func Add() gin.HandlerFunc{
	return func(c *gin.Context){
		var model models.ProductItem
		err := c.ShouldBindJSON(&model)

		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error": err,
			})
		}else{
			err := model.AddProduct()
			if err != nil{
				c.JSON(http.StatusInternalServerError,gin.H{
					"error": err,
				})
			} else{
				c.JSON(200,gin.H{
					"success" : "addedd successfully",
				})
			}
		}
	}
}