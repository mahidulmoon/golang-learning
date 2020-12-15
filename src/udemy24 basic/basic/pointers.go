package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	p := person{
		name: "mahidul",
		age:23,
	}
	p.changeName("moon")
	fmt.Println("without pointer:",p)

	p.changeAge(25)
	fmt.Println("wit pointer:",p)
}

func (p person) changeName(name string) {
	p.name = name
}
func (p *person) changeAge(age int) {
	p.age = age
}