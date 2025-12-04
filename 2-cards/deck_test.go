package main

import (
	"errors"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if (len(d) != 52) {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if (d[0] != "Ace of Spades") {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	if (d[len(d)-1] != "King of Diamonds") {
		t.Errorf("Expected first card to be King of Diamonds but got %v", d[0])
	}
}

func TestSaveFileAndTestLoadFile(t *testing.T) {
	// Delete old files first
	err := os.Remove("_decktesting")

	if (err != nil && !errors.Is(err, os.ErrNotExist)) {
		t.Error(err)
	} 

	d := newDeck()
	d.save("_decktesting")
	deck := read("_decktesting")

	if (len(deck) != 52) {
		t.Error("Error in reading file")
	}

	err = os.Remove("_decktesting")

	if (err != nil) {
		t.Error("Unable to delete testing files after testing")
	} 
}