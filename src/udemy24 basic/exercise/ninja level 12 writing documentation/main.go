package main

import (
	"fmt"
	"./dog"
)

type canine struct {
	name string
	age int
}

func main(){
	fudo := canine{
		name: "Fidu",
		age: dog.Years(2),
	}
	fmt.Println(fudo)
}