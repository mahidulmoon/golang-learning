package main

import "fmt"

const PI float32 = 3.1416

const (
	name = "moon"
	//age int //will show error
)

func main() {
	fmt.Println("Constant assign outside function:",PI)
	//age = 23 //cannot assign here
	fmt.Println("Constant assign outside function:",name)
	//fmt.Println("Constant assign inside function:",age) //will show error

}
