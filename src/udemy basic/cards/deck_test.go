package main

import (
	"os"
	"testing"
)

func newDeckTest(t *testing.T) {
	newdeck := newDeck()

	if len(newdeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v",len(newdeck))
	}
}

func TestDecktosaveandTakeNewDeckFromFile (t *testing.T){
	os.Remove("_decTesting")

	deck := newDeck()

	deck.savetoFile("_decTesting")

	loadDeck := newDecFromFile("_decTesting")

	if len(loadDeck) != 16{
		t.Errorf("Expecting deck length is 16 but got %v",len(loadDeck))
	}

	os.Remove("_decTesting")
}