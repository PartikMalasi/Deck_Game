package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Acegam of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
}
