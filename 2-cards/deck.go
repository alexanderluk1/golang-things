package main

import "fmt"

// Create a new type of 'Deck' which is a slice of strings which extends string
type deck []string

func newDeck() deck {
	cardSuite := []string{"Spade", "Heart", "Club", "Diamond"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var cards []string

	for _, val := range cardValue {
		for _, suit := range cardSuite {
			combined := val + " of " + suit
			cards = append(cards, combined)
		}
	}
	return cards
}

// Loop through deck of cards and print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}