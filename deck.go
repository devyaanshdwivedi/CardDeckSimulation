package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // reading the file
	if err != nil {
		fmt.Println("Error:", err) // printing the error if any
		os.Exit(1)                 // exiting the program with an error code
	}

	//string(bs) // converting the byte slice to a string
	s := strings.Split(string(bs), ",") // splitting the string into a slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource((time.Now().UnixNano())) // creating a new source of randomness
	r := rand.New(source)                             // creating a new random number generator

	for i := range d { // iterating over the deck of cards
		newPosition := r.Intn(len(d) - 1)           // generating a new random position
		d[i], d[newPosition] = d[newPosition], d[i] // swapping the cards
	}
}
