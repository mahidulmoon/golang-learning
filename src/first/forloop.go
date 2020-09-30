package main
import "fmt"

func main(){
	for i:=0;i<=10;i++{
		fmt.Println(i)
	}
	// name := "moon"
	// for _,char:= range(name){
	// 	fmt.Println(string(char))
	// }
	
	a := 0

	for a<5{
		fmt.Println("moon")
		a++
		if a==3{
			break
		}
	}
}