package main

import "fmt"

const (
	x = 10
	y = 50
)

func main()  {
	a := (42 == 43)
	b := (42 >= 43)
	c := (42 <= 43)
	d := (42 != 43)
	e := (42 > 3)
	f := (42 < 43)

	fmt.Println(a," + ",b," + ",c," + ",d," + ",e," + ",f)

	fmt.Println("/n")
	fmt.Println("untypes const:",x," + ",y)

	variable := `this is raw string 
					which is written by md "mahidulmoon"`
	fmt.Println(variable)
}
