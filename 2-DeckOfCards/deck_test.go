package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(cards))
	}

	if cards[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be the Ace of Hearts but got %v", cards[0])
	}

	if cards[51] != "King of Clubs" {
		t.Errorf("Expected final card to be the King of Clubs but got %v", cards[51])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(loadedDeck))
	}

}
