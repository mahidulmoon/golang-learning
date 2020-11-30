package main

import "fmt"

type contactInfo struct{
	email string
	zip int
}

type person struct{
	firstname string
	lastname string
	contact contactInfo
}


func main()  {
	name := person{firstname: "mahidul",lastname: "moon",contact:contactInfo{email:"mahidulmoon@gmail.com",zip:2216},}
	fmt.Println(name)

	var updatename person

	updatename.firstname = "moon"
	updatename.lastname = "mahidul"

	fmt.Println(updatename)
}