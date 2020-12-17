package main

import "fmt"

func main(){
	mySlice := []string{"m","0","o","n"}

	fmt.Println(mySlice)
	fmt.Println(mySlice[0:2])
	fmt.Println(mySlice[3])

	mySlice[3] = "name"

	fmt.Println(mySlice)

	// delete(mySlice,"name") // cannot use delete like this
	//fmt.Println(mySlice)
}
