package main

func main() {
	cards := deck{newCard(), "Ace of Spades"}

	cards.print()
}

func newCard() string {
	return "Four of Diamonds"
}