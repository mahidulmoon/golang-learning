package main

import "fmt"

func main() {
	i := 12

	if i < 5 {
		fmt.Println("true")
	}else if i > 10{
		fmt.Println("else if true")
	}else {
		fmt.Println("else true")
	}
}
