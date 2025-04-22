package main

//import "fmt"

func main() {
	cards:= newDeck()
	cards.shuffle() // shuffle the deck of cards
	cards.print() // print the deck of cards
}
