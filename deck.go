package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// do not add a receiver, only add when we want to call a method name, i.e. dot method name

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// any var of type deck now gets access to the print method
// d is the actual copy of the deck we are working with is available in the function as a variable called 'd'
// every var of type 'deck' can call this fcn on itself
// refer to the receiver (d) by a 1 or 2 letter abbreviation that matches the type of the receiver

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
