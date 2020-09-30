package main
import "fmt"

func main(){
	functionname()
	funtiontoadd(3,4,"moon")
	result:= functionreturn(5)
	fmt.Println(result)
	
}

func functionname(){
	fmt.Println("hello")
}
func funtiontoadd(a int,b int,c string){
	fmt.Println(a+b,c)
}

func functionreturn(x int) int{
	Result:= x*5/3
	return Result
}