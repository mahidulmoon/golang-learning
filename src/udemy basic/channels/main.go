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
		//go responseCheck(link) // go only use before function calling and it is not care about waiting for child function response just execute
	}
}


func responseCheck(link string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link,"Might be down")
	}
	fmt.Println(link,"server is ok")
}