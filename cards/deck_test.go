package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 24 {
		t.Errorf("Expected length 24, but got %d", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts but got %v", d[0])
	}

	if d[len(d)-1] != "Six of Clubs" {
		t.Errorf("Expected first card of Six of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 24 {
		t.Errorf("Expected length 24, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
