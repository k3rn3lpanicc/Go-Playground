package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Error("First element is not correct")
	}
	if d[len(d)-1] != "king of Clubs" {
		t.Error("Last element is not correct")
	}
}
