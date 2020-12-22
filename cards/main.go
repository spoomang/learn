package main

func main() {
	cards := NewDeck()

	hand, _ := deal(cards, 5)

	hand.toString()
}
