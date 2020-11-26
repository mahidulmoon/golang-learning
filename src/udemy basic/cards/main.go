package main

import "fmt"

func main() {
	//var card string = "Titanic" //scope variable

	card := "Rise of Empire"
	card = "Using old variable" //if variable is declared before than no need to give : befor =

	fmt.Println(card)

	card_from_func := newCard()
	fmt.Println(card_from_func)
}

func newCard() string {
	return "Game of thrones"
}
