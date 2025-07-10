package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	newCards := newDeck()
	newCards.saveToFile("_decktesting")

	loadedCards := newDeckFromFile("_decktesting")

	if len(loadedCards) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedCards))
	}

	os.Remove("_decktesting")
}
