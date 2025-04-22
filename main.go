package main

//import "fmt"

func main() {

	cards := newDeck()            // calling the function to assign value to cards
	cards.saveToFile("my_cards") // calling the function to save the deck to a file
}
