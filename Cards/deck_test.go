package main

import (
	"os"
	"testing"
)

// func TestNewCard(t *testing.T) {
// 	d := newDeck()
// 	d.Shuffle()
// 	if len(d) != 16 {
// 		t.Errorf("Expected 16 but recieved %v", len(d))
// 	}

// 	if d[0] != "Ace of Spades" {
// 		t.Errorf("Expexted first element as Ace of Spades, but got %v", d[0])
// 	}
// 	if d[len(d)-1] != "Four of Clubs" {
// 		t.Errorf("Expexted last element as Four of Clubs, but got %v", d[len(d)-1])
// 	}
// }

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 but recieved %v", len(d))
	}

	os.Remove("_decktesting")
}
