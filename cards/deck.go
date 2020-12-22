package main

import (
	"fmt"
	"strings"
)

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spade", "Diamond", "Hearts", "Clauver"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d Deck) print() {
	for _, value := range d {
		fmt.Println(value)
	}
}

func (d Deck) toString() {
	s := strings.Join(d, "|")
	fmt.Println(s)
}

func deal(deck Deck, hand int) (d1 Deck, d2 Deck) {
	d1 = deck[:hand]
	d2 = deck[hand:]
	return
}
