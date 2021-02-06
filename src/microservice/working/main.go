package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(rw http.ResponseWriter,r*http.Request){
		log.Println("Hello world")
		d,_ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s",d)
	})
	http.HandleFunc("/goodbye",func(http.ResponseWriter,*http.Request){
		log.Println("Good Bye world")
	})

	http.ListenAndServe(":9090",nil)
}
