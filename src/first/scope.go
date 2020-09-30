package main
import "fmt"

var moon int = 34 //global only way declaration


func main(){
	mahidul := 22 //works only inside function

	fmt.Println(moon)
	fmt.Println(mahidul)
	
}

func localfunction(){
	fmt.Println(moon)
	// fmt.Println(mahidul)
}