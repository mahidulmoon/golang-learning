package main

import "fmt"

func main(){
	c := make(chan int)
	//c <- 42		// this is not going to work without anonymous(lamda) function

	go func() {
		c <- 42
	}()


	fmt.Println(<-c) //going to block the channel for a while
}
