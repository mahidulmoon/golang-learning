package main

import "fmt"

type person struct {
	firstname string
	lastname string
	favFlavours []string
	body body
}

type body struct {
	color string
	size string
}

func main(){
	p := person{
		firstname: "mahidul",
		lastname: "moon",
		favFlavours: []string{"cream","slice","chocolate"},
		body: body{
			color:"red",
			size:"fat",
		},
	}
	fmt.Println(p)

	for _,v := range p.favFlavours{
		fmt.Println(v)
	}

	m := map[string] person{
		p.lastname:p,
	}
	fmt.Println(m)

	anonym := struct {
		name string
		age int
	}{
		name:"moon",
		age: 23,
	}
	fmt.Println(anonym)
}