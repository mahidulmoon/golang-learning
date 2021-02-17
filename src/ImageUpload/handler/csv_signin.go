package handlers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type SiteRegister struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

func UploadCSV() gin.HandlerFunc{
	return func(c *gin.Context){
		csvfile,_,err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}else{
			r := csv.NewReader(csvfile)
			for {

				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				//fmt.Printf("Name: %s Email: %s Phone: %s Password: %s\n", record[0], record[1],record[2],record[3])
				user := SiteRegister{
					Name: record[0],
					Email: record[1],
					Phone: record[2],
					Password: record[3],
				}
				bs,err := json.Marshal(user)
				if err != nil{
					fmt.Println("error to marshal csv input")
				}
				//user_json := string(bs)
				//fmt.Println(user_json) //json format data

				resp,_ := http.Post("http://127.0.0.1:8083/api/validate","application/json",bytes.NewBuffer(bs))

				defer resp.Body.Close()

				//body,_ := ioutil.ReadAll(resp.Body)
				//
				//
				//fmt.Println(string(body))

				if ( resp.StatusCode == 200 ){
					fmt.Println("email and phone number found")
				}else{
					//fmt.Println("signup")
					register_resp,_ := http.Post("http://127.0.0.1:8083/api/register/","application/json",bytes.NewBuffer(bs))
					defer register_resp.Body.Close()

					c.JSON(http.StatusOK,gin.H{
						"message" : "csv file process successful",
					})
				}



			}
			return
		}
	}
}