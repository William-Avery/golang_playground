package main

import (
	"os"
	"testing"
)

func TestDeckSize(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestSaveandLoadDeck(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := LoadDeck("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in Deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
