package main

import (
	"github.com/mahidulmoon/banking/handler"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){

	//mux := http.NewServeMux()
	mux := mux.NewRouter()

	mux.HandleFunc("/greet", handler.Greet)
	mux.HandleFunc("/customers",handler.GetAllCustomer)

	http.ListenAndServe("localhost:8000",mux)
}
