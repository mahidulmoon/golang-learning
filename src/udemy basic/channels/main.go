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
	c := make(chan string)

	for _,link := range links{
		responseCheck(link,c)
		//go responseCheck(link) // go only use before function calling and it is not care about waiting for child function response just execute
	}
	fmt.Println(<-c)
}


func responseCheck(link string,c chan string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link,"Might be down")
		c <- "Might be down I think"
	}
	fmt.Println(link,"server is ok")
	c <- "Yes It is fine"

}