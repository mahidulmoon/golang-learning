package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	fmt.Println("Begin CPU:",runtime.NumCPU())
	fmt.Println("Begin GS:",runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		fmt.Println("Hello 1")
		wg.Done()
	}()
	go func(){
		fmt.Println("Hello 2")
		wg.Done()
	}()

	fmt.Println("Mid CPU:",runtime.NumCPU())
	fmt.Println("Mid GS:",runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Main run successfully")

	fmt.Println("END CPU:",runtime.NumCPU())
	fmt.Println("END GS:",runtime.NumGoroutine())
}
