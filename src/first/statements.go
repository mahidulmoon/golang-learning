package main
import "fmt"

func main(){
	moon := 8>9

	if moon{
		fmt.Println(moon)
	}else{
		fmt.Println(moon)
	}

	if moon && 3>2{
		fmt.Println("Working")
	}else{
		fmt.Println("Not working")
	}
	if moon || 3>2{
		fmt.Println("Working")
	}else{
		fmt.Println("Not working")
	}
}