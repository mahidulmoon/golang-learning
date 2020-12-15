package main

import "fmt"

var x int = 40
var y int

func main(){
	fmt.Println("this is inside main function.")
	newFunc()
	getValue := varChange()
	fmt.Println("Getting value from function: ",getValue)
}

func newFunc() {
	fmt.Println("Running inside another function")
	fmt.Println(x)
}

func varChange() int{
	y = 100
	return y
}