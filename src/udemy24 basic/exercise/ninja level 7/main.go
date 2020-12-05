package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	p := person{}
	p.changeName("moon")
	fmt.Println(p)
	fmt.Println(p.changeAge())
}

func (p *person) changeName(name string){
	p.name = name
}

func (p *person) changeAge() int{
	//fmt.Println(&p)  //this will show the memory address
	p.age=23
	return p.age
}