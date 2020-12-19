package main

import "fmt"

func main() {
	n := hashBucket("Go",12)
	fmt.Println(n)
}

func hashBucket(value string,number int) int {
	letter := int(value[0])
	//fmt.Println(letter)
	bucket := letter % number
	return bucket
}