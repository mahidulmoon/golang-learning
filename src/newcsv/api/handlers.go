package api

import (
	"encoding/csv"
	"fmt"
	"github.com/Upskillbd/newcsv/models"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func CreateHandler()gin.HandlerFunc{
	return func(c *gin.Context){

		var values models.QueryStruct
		filename := c.Param("filename")
		err := c.ShouldBindJSON(&values)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "could not bind json",
			})
		}else{
			var list []string
			file, err := os.OpenFile("../files/test"+filename+".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			result := models.GetRows(values.Query)
			//fmt.Println(reflect.TypeOf(result))
			//json,_ := os.Stdout.Write(result)

			//fmt.Println(string(result))
			st := strings.Fields(strings.Trim(fmt.Sprint(string(result)), "[{}]"))
			wr := csv.NewWriter(file)
			//m := make(map[string]string)
			//err := json.Unmarshal(result, &m)
			//fmt.Println(reflect.TypeOf(st))

			fmt.Println(err)

			for _,row := range st{
				if row != "},"{
					if row != "{" {
						list = append(list,row)
					}
				}else{
					//fmt.Println("break")
					wr.Write(list)
					list = nil
				}
				//fmt.Println(row)
			}
			wr.Flush()

			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				//"result" : result,
				"link" : "http://127.0.0.1:9004/api/download/"+filename,
			})
		}

	}
}

func DownloadAttachmentHandler()gin.HandlerFunc{
	return func(c *gin.Context){
		filename := c.Param("filename")

		file, err := os.Open("../files/test"+filename+".csv") //Create a file
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