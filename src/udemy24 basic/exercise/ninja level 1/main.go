package main

import "fmt"

var a int
var b string
var c bool


type T1 int
type T2 string
type T3 bool

var m T1
var n T2
var o T3

func main(){
	x,y,z := 42,"moon",true
	fmt.Println(x," + ",y," + ",z)

	fmt.Println("Variable declare from scope")
	fmt.Println(a," + ",b," + ",c)

	fmt.Println("Assigning value into scope variable")
	var a,b,c = 219,"mahidul",true
	fmt.Println(a," + ",b," + ",c)

	//redeclaring old variable
	b = "mahidulmoon"
	fmt.Println(b)

	fmt.Println("Assigning value type variable")
	var m,n,o = 230,"runa",false
	fmt.Println(m," + ",n," + ",o)
}
