package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of 'deck' which is a slice of strings
type deck []string

//create a new function which belongs to the deck type
//print each card in the deck

func newDeck() deck {
	cards := deck{} // creating a new deck of cards

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // appending the card to the deck
		}
	}
	return cards // returning the deck of cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card) // printing the index and the card
		
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // returning the hand and the remaining deck
}

func (d deck) toString() string {
	return strings.Join(d, ",") // joining the deck of cards into a string
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // writing the deck to a file
}