package main

import "fmt"

func main(){
	defer fmt.Println("this is from main function")

	func(){
		fmt.Println("this is inside self executing anonymous")
	}()
}
