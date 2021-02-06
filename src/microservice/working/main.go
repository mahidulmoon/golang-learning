package main

import (
	"log"
	"microservice/working/handlers"
	"net/http"
	"os"
)

func main(){
	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	hh := handlers.NewHello(l)
	//http.HandleFunc("/goodbye",func(http.ResponseWriter,*http.Request){
	//	log.Println("Good Bye world")
	//})
	sm := http.NewServeMux()
	sm.Handle("/",hh)
	http.ListenAndServe(":9090",sm)
}
