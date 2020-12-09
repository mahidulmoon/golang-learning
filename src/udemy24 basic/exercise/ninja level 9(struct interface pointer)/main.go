package main

import "fmt"

type person struct {
	firstname string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Human is talking")
}

func saySomething(h human){
	h.speak()
}


func main(){

	p1 := person{}
	saySomething(&p1) // if function receive pointer receiver than param must be declared with pointer before pass
}