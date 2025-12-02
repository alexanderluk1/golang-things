package main

import "fmt"

// Create a new type of 'Deck' which is a slice of strings which extends string
type deck []string

// Loop through deck of cards and print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}