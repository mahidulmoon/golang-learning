package main

import "fmt"

func main(){
	colors := map[string] string{
		"red" : "#ff000",
		"green" : "$4bf745",
	}

	// var empty map[string]string     //empty map declared
	// empty := make(map[string]string)     //empty map declared with builtin function

	colors["white"] = "#ffffff"   //adding new value in map with key

	fmt.Println(colors)
}
