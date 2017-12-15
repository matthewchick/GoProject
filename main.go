package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades" // static type
	card := newCard() //initialize
	// card = "Five of Diamonds"   //reassignment

	fmt.Println(card)
}

func newCard() string { //return a string
	return "Five of Diamonds"
}
