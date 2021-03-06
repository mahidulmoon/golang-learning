package main

import (
	"github.com/mahidulmoon/banking/handler"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){

	//mux := http.NewServeMux()
	mux := mux.NewRouter()

	mux.HandleFunc("/greet", handler.Greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers",handler.GetAllCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customer",handler.CreateCustomer).Methods(http.MethodPost)
	mux.HandleFunc("/customer/{customer_id}",handler.ResponseFromUrl).Methods(http.MethodPost)

	http.ListenAndServe("localhost:8000",mux)
}
