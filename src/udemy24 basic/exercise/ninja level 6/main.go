package main

import "fmt"

type person struct {
	firstname string
	lastname string
	age int
}

func main(){
	number := foo()
	dob,dob_s := bar()

	defer fmt.Println(dob_s," : ",dob,"and my id is :",number)
	list := []int {1,2,3,4,5}
	result := foo_param(list...)
	fmt.Println("Result: ",result)

	p := person{
		firstname: "mahidul",
		lastname: "moon",
		age:23,
	}
	p.structShow()
	func (){
		fmt.Println("Inside anonymus function")
		for i:=0;i<5;i++{
			fmt.Print(i," ")
		}
		fmt.Println()
	}()
}

func foo() int{
	return 219
}

func bar () (int,string){
	return 1997,"date of birth"
}

func foo_param(xi ...int) int {
	total := 0
	for _,v := range xi{
		total+=v
	}
	return total
}

func (p person)structShow(){
	fmt.Println("hello '",p.firstname,"' your nickname is '",p.lastname,"' and your age is ",p.age)
}