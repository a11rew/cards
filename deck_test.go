package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Validate 56 cards are output
	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, got: %v", d[0])
	}

	if d[len(d)-1] != "One of Clubs" {
		t.Errorf("Expected last card to be One of Clubs, got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	// Delete any existing temp files
	const tfn = "_decktesting"

	os.Remove(tfn)

	d := newDeck()

	d.saveToFile(tfn)

	loadedDeck := newDeckFromFile(tfn)

	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, got: %v", len(d))
	}

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected saved deck length to match recovered deck length")
	}

	os.Remove(tfn)
}
