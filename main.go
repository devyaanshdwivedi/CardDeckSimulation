package main

//import "fmt"

func main() {
	//cards := []string{"Ace of Diamonds", newCard()}
	cards:= newDeck() // calling the function to assign value to cards

	hand, remainingCards := deal(cards, 5) // calling the function to deal cards

	hand.print()
	remainingCards.print() // printing the remaining cards
}
	
