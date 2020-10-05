package handlers

import(
	"github.com/gin-gonic/gin"
	"practicefolder/models"
	"net/http"
	"strconv"
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

func GetProductBytid() gin.HandlerFunc{
	return func(c *gin.Context){
		ID,_ := strconv.Atoi(c.Param("id"))
		model := models.ProductItem{}
		
		model,err := models.GetProductsById(ID)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"error": err,
			})
		}else{
			c.JSON(http.StatusOK,model)
		}
	}
}

func DeleteByID() gin.HandlerFunc{
	return func(c *gin.Context){
		ID,_ := strconv.Atoi(c.Param("id"))
		//model := models.ProductItem{}
		
		err := models.DeleteProductsById(ID)
		if err != nil{
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"erro": err,
			})
		}else{
			c.JSON(200,gin.H{
				"success":"deleted successfully",
			})
		}
	}
}