package main

import "fmt"

func main(){
	fmt.Println("Give the maximum numer : ")
	var max_num int
	fmt.Scanln(&max_num)
	number := 0
	for i:=0;i<=max_num;i++{
		if i%2 == 0 {
			fmt.Print(i%2," ")
			number++
		}
	}
	fmt.Println()
	fmt.Println("Number of Zero",number)

	fmt.Println("Insert id from (219/230):")
	var id int
	fmt.Scanln(&id)
	switch id {
	case 219:
		fmt.Println("mahidulmoon")
	case 230:
		fmt.Println("runa")
	}

}
