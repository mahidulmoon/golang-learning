package main

import (
	"encoding/json"
	"fmt"
)

//to convert into json struct value must be capitalized
type Human struct {
	First string
	Last string
	Sayings []string
}

func main(){
	p := Human{
		First: "mahidul",
		Last: "moon",
		Sayings: []string{"Hello","Good morning","fuck you"},
	}
	bs,err := json.Marshal(p)

	if err != nil{
		fmt.Println("error to marshal")
	}
	fmt.Println(string(bs)) // final conversion to json
}

