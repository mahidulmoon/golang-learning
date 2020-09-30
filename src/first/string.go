package main
import "fmt"

func main(){
	first := "mahidul"
	last := "moon"

	fmt.Println(first+last)
	totalstring := first+last
	fmt.Println( totalstring[1:4] )
	fmt.Println( totalstring[5] )
	fmt.Println( totalstring[1:] )
	fmt.Println( totalstring[:4] )
	
}