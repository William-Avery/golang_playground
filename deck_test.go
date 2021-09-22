package main

import "testing"

func TestDeckSize(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
