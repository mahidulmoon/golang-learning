package main
import "fmt"

func main(){
	moon := "good"

	switch (moon){
	case "verygood":
		fmt.Println("case1")
	
	case "good":
		fmt.Println("case2")
	
	case "bad":
		fmt.Println("case3")
	
	default:
		fmt.Println("this is default")
	}
	
}