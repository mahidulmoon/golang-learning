package main

import (
	"fmt"
db	"pg_tuts/db"
)


func main(){
	fmt.Println("hello")
	db.Connect()
}