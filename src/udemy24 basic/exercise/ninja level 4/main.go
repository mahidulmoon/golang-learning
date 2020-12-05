package main

import "fmt"

func main(){
	arr := []int{1,2,3,4,5}

	for i:=0;i<5;i++{
		arr = append(arr,i)
	}
	fmt.Println(arr)
	fmt.Println(arr[5:])
	arr1:=[]string{"hello","mahidul","moon","how","are you"}
	arr2 := []string{"hello","mahidul","moon","how","are you",}

	xx2 := [][]string{arr1,arr2}
	fmt.Println(xx2)

	for i,xs := range xx2{
		fmt.Println("index:",i)
		for j,v := range xs{
			fmt.Println("Index:",j,"Value:",v)
		}
	}
}
