package main

import "fmt"

func main(){
	c := generate()
	recieve(c)
	fmt.Println("going to exit")
}

func recieve(c <-chan int){
	for v:= range c{
		fmt.Println(v)
	}
}

func generate() <-chan int{
	c := make(chan int)

	go func() {
		for i:=0;i<10;i++{
			c <- i
		}
		close(c)
	}()

	return c
}