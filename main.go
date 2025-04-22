package main

import "fmt"

//import "fmt"

func main() {

	cards := newDeck()            // calling the function to assign value to cards
	fmt.Println(cards.toString()) // calling the function to convert the deck to a string
}
