// Every element in a slice must be of the same type which is the same as array

package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades" // static type
	cards := deck{"Ace of Diamonds", newCard()} //initializeby by a slice of type string

	// card = "Five of Diamonds"   //reassignment
	cards = append(cards, "Six of Spades") //not modify the slice

	cards.print() //call print function from deck.go

	for _, card := range cards {
		fmt.Println(card)
	}

}

func newCard() string { //return a string
	return "Five of Diamonds"
}
