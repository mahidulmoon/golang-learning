package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main(){
	p1:=person{first: "mahidul",last: "moon",age: 23}
	p2:=person{first: "meher",last: "afroz",age: 23}

	fmt.Println(p1)
	p2.changeName("Md. Meher")
	fmt.Println(p2)
}

func(p *person) changeName(name string){
	p.first = name
}