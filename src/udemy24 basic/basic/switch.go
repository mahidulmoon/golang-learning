package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter a nickname:")
	fmt.Scan(&name)
	switch name {
	case "moon":
		fmt.Println("mahidulmoon")
	case "runa":
		fmt.Println("rahima akter")
	default:
		fmt.Println("nothing matched")

	}

}
