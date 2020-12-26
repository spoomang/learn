package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	expectedNumber := 16
	assert.Equal(t, expectedNumber, len(deck))
	assert.Equal(t, "Ace of Spade", deck[0])
	assert.Equal(t, "Ace of Diamond", deck[4])
	assert.Equal(t, "Ace of Hearts", deck[8])
	assert.Equal(t, "Ace of Clauver", deck[12])
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := NewDeck()
	deck.saveToFile("_deckTesting")

	d := newDeckFromFile("_deckTesting")

	expectedNumber := 16
	assert.Equal(t, expectedNumber, len(d))
	assert.Equal(t, "Ace of Spade", d[0])
	assert.Equal(t, "Ace of Diamond", d[4])
	assert.Equal(t, "Ace of Hearts", d[8])
	assert.Equal(t, "Ace of Clauver", d[12])

	os.Remove("_deckTesting")
}
