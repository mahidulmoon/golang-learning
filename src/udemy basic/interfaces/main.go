package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct {}
type spanishBot struct {}


func main(){
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings()string{
	return "Hello there!"
}

func (spanishBot) getGreetings() string{
	return "Hola"
}