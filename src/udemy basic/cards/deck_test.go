package main

import (
	"fmt"
	"testing"
)

func newDeckTest(t *testing.T) {
	newdeck := newDeck()

	if len(newdeck) != 16 {
		fmt.Println("Expected dec length was 20, but got ", len(newdeck))
	}
}
