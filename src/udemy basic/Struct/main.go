package main

import "fmt"

type person struct{
	firstname string
	lastname string
}

func main()  {
	name := person{firstname: "mahidul",lastname: "moon"}
	fmt.Println(name)

	var updatename person

	updatename.firstname = "moon"
	updatename.lastname = "mahidul"

	fmt.Println(updatename)
}