package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'Deck' which is a slice of strings which extends string
type deck []string

func newDeck() deck {
	cardSuite := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
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

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) save(location string) error {
	return os.WriteFile(location, []byte(d.toString()), 0666)
} 

func read(path string) (deck) {
	bs, err := os.ReadFile(path)

	if (err != nil) {
		log.Fatal(err)
		os.Exit(1)
	}
	convertedStr := string(bs)

	return deck(strings.Split(convertedStr, ","))
}

func (d deck) shuffle() {
	for i := range d {
		randIndex := rand.Intn(len(d)-1)
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}