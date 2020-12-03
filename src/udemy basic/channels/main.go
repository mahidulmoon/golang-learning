package main

import (
	"fmt"
	"net/http"
)

func main(){
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _,link := range links{
		responseCheck(link)
	}
}


func responseCheck(link string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link,"Might be down")
	}
	fmt.Println(link,"server is ok")
}