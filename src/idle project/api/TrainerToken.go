package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainerList struct {
	Url string
}

var Instance = TrainerList{"https://service.upskill.com.bd/public"}

type UpskillTrainer struct {
	ID string `json:"id"`
}

func (o *TrainerList) GetTrainerURL() string {
	return o.Url + "/cons"
}

func (o *TrainerList) GetTrainerInfoURL() string {
	return o.Url + "/consultant/info"
}

func (o *TrainerList) GetUserFromToken(token string) (UpskillTrainer, error) {
	user := UpskillTrainer{}
	err := GetJsonAndBindAuth(o.GetTrainerInfoURL(), token, &user)
	return user, err
}

//from util
func GetJsonAndBindAuth(url string, token string, obj interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Set appropriate headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Secret", token)

	//fmt.Println("2) Token from GETJSONBIND", token)
	// Prepare an HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	jsonBody, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("3) GetJsonBind", string(jsonBody))
	err = json.Unmarshal(jsonBody, &obj)
	if err != nil {
		return err
	}
	return nil
	//if err != nil {
	//	log.Printf("error decoding sakura response: %v", err)
	//	if e, ok := err.(*json.SyntaxError); ok {
	//		log.Printf("syntax error at byte offset %d", e.Offset)
	//	}
	//	log.Printf("sakura response: %q", resp.Body)
	//	return err
	//}
	//fmt.Println(resp.Body)
	//return nil
}

//--------example
func TrainerTokenCatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Auth-Secret")
		TrainerID, err := Instance.GetUserFromToken(token)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
		c.JSON(http.StatusFound, gin.H{
			"userinfo": TrainerID.ID,
		})

	}
}
