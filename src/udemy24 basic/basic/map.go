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
}
