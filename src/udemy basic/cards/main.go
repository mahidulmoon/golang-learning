package main

import "fmt"

func main() {
	//var card string = "Titanic" //scope variable

	card := "Rise of Empire"
	card = "Using old variable" //if variable is declared before than no need to give : befor =

	fmt.Println(card)

	card_array := deck{"Heart of diamonds", newCard()} //for deck go run main.go deck.go

	card_array = append(card_array, "six of spades")

	card_array.print() //calling function from another file

	card_from_newDeck := newDeck()
	card_from_newDeck.print()

	//returning multiple value and receiving that
	hand, remainingCards := deal(card_from_newDeck, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of diamonds"
}
