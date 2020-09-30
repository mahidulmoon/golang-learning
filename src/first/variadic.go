package main
import "fmt"

func main(){
	result:=add(3,4,5,6,7,8,9)
	fmt.Println(result)
}

func add(nums ...int)int{
	var totale int
	for _,n:= range nums{
		totale +=n 
	}
	return totale
}