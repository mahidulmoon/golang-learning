package main

import "fmt"

var global_name string = "mahidulmoon"




func main(){
	a := 3
	b := "moon"
	c := 4.3
	d := true

	m,n := 6,7

	fmt.Printf("%v \n",a)
	fmt.Printf("%v \n",b)
	fmt.Printf("%v \n",c)
	fmt.Printf("%v \n",d)
	fmt.Printf("%v \n",global_name)
	fmt.Printf("%v \n",m)
	fmt.Printf("%v \n",n)

	//existing variable value assign
	b = "mahidul moon"
	fmt.Printf("%v \n",b)
}
