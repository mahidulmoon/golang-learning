package api

import (
	"encoding/csv"
	"io"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"github.com/UpskillBD/gopgtocsv/models"

)

func GETUser()gin.HandlerFunc{
	return func (c *gin.Context){
		user, err := models.GetUser()
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err,
				"message" : "error found",
			})

		} else {
			var list []string
			file, err := os.OpenFile("../files/test.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			st := strings.Fields(strings.Trim(fmt.Sprint(user), "[]"))
			wr := csv.NewWriter(file)

			for i,value := range st{
				if value == "{{}"{
					st[i] = "\n"
				}
			}
			//fmt.Println(st)
			for _,value := range st{
				if value != "\n"{
					list = append(list,value)
				}else{
					//fmt.Println("break")
					wr.Write(list)
					list = nil
				}
			}

			wr.Flush()



			c.JSON(http.StatusInternalServerError,gin.H{
				"message" : "conversion done",
				"link" : "http://127.0.0.1:9004/api/download",
			})

		}
	}
}

func DownloadAttachmentHandler()gin.HandlerFunc{
	return func(c *gin.Context){


		file, err := os.Open("../files/test.csv") //Create a file
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": "Error:" + err.Error(),
			})
			return
		}
		defer file.Close()
		c.Writer.Header().Add("Content-type", "text/csv")
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": "Error:" + err.Error(),
			})
			return
		}
	}
}