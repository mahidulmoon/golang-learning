package main

import "fmt"

func main() {
	x := make(map[string]int)

	x["age"] = 23
	x["number"] = 98

	fmt.Println(x)

	x["number"] = 10
	fmt.Println(x)

	delete(x,"age")
	fmt.Println(x)

	newMap:= map[int]string{
		0:"hey",
		1:"mahidul",
		2:"moon",
	}

	for key,value := range newMap{
		fmt.Println("Key:",key,"and value is :",value)
	}
}
