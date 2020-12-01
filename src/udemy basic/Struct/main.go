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

	updatename.print()

	name.newUpdate("islam")
	name.print()
}

func (p person) print(){
	fmt.Println(p)
}

func (p *person) newUpdate(newName string){  //without pointer go just pass the copy of the struct
	p.lastname = newName
}

// &variable  -> give the memory address of the value is this variable pointing at
// *pointer   -> give me the value of this memory address is pointing at