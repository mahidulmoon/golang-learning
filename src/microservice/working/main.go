package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(rw http.ResponseWriter,r*http.Request){
		log.Println("Hello world")
		d,err := ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw,"Oops",http.StatusBadRequest)
		}
		fmt.Printf("Data %s",d)
	})
	http.HandleFunc("/goodbye",func(http.ResponseWriter,*http.Request){
		log.Println("Good Bye world")
	})

	http.ListenAndServe(":9090",nil)
}
