package main

// import "fmt"

func main() {
	cards := newDeck()
	
	// fmt.Println(cards.toString())

	hand, remaining := cards.deal(5)

	hand.save("hand")
	remaining.save("deck")

	savedHand := read("hand")
	savedRemaining := read("deck")

	savedHand.print()
	savedRemaining.print()

	// hand, remaining := cards.deal(5)

	// hand.print()
	// fmt.Println("---")
	// remaining.print()
}