package handler

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string  `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}


func GetAllCustomer(w http.ResponseWriter,r *http.Request){
	cutomer := []Customer{
		{"Moon","Dhaka","3213"},
		{"Mahidul","Chittagong","12345"},

	}
	//check the header input and present data according to input
	if r.Header.Get("Content-Type") == "application/json"{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(cutomer)
	}else{
		w.Header().Add("Content-Type","application/xml")
		xml.NewEncoder(w).Encode(cutomer)
	}
}