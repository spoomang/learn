package main

func main() {
	cards := NewDeck()
	cards.shuffle()
	cards.print()
}
