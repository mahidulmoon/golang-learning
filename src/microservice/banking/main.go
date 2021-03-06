package main

import (
	"github.com/mahidulmoon/banking/handler"
	"net/http"
)

func main(){
	http.HandleFunc("/greet", handler.Greet)
	http.HandleFunc("/customers",handler.GetAllCustomer)

	http.ListenAndServe("localhost:8000",nil)
}
