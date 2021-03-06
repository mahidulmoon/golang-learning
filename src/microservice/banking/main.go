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

	http.ListenAndServe("localhost:8000",mux)
}
