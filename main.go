package main

//import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print() // printing the deck of cards
}
