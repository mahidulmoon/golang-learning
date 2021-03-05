package handler

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name string
	City string
	Zipcode string
}


func GetAllCustomer(w http.ResponseWriter,r *http.Request){
	cutomer := []Customer{
		{"Moon","Dhaka","3213"},
		{"Mahidul","Chittagong","12345"},

	}
	json.NewEncoder(w).Encode(cutomer)
}