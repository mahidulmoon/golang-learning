package main

import (

	"testing"
)

func newDeckTest(t *testing.T) {
	newdeck := newDeck()

	if len(newdeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %v",len(newdeck))
	}
}
