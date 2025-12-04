package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	hand, remaining := cards.deal(5)

	hand.print()
	fmt.Println("---")
	remaining.print()
}