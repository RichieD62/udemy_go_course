package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}

// New Deck - create a new list of playing cards which is essentially an array of strings
// Print - Log out the contents of a deck of cards
// Shuffle - shuffle all the cards in the deck
// Deal - Create a hand of cards
// Save To File - Save a list of cards to a file on the local machine
// New Deck From File - Load a list of cards from the local machine
