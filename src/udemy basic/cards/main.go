package main

import "fmt"

func main() {
	//var card string = "Titanic" //scope variable

	card := "Rise of Empire"
	card = "Using old variable" //if variable is declared before than no need to give : befor =

	fmt.Println(card)

	card_array := deck{"Heart of diamonds", newCard()} //for deck go run main.go deck.go

	card_array = append(card_array, "six of spades")

	for i, cards := range card_array {
		fmt.Println(i, cards)
	}
}

func newCard() string {
	return "Five of diamonds"
}