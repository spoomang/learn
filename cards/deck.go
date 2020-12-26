package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck array of strings
type Deck []string

// NewDeck creates a new deck
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
	for i, value := range d {
		fmt.Printf("%d %v\n", i, value)
	}
}

func (d Deck) toString() string {
	s := strings.Join(d, "|")
	return s
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func deal(deck Deck, hand int) (d1 Deck, d2 Deck) {
	d1 = deck[:hand]
	d2 = deck[hand:]
	return
}

func newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "|")

	return Deck(s)
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
