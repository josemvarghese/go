package main

import (
	"fmt"
)

func main() {
	var card string = newCard()
	var cardBool bool = true
	var cardInt int = 123
	var cardFloat float32 = 342.23
	// Static type language
	fmt.Println("Cards part 1")
	fmt.Println(card)
	fmt.Println(cardBool)
	fmt.Println(cardInt)
	fmt.Println(cardFloat)

	// alternative for above
	cardalt := newCard()
	cardBoolalt := false
	cardIntalt := 321
	cardFloatalt := 5642.23
	fmt.Println(cardalt)
	fmt.Println(cardBoolalt)
	fmt.Println(cardIntalt)
	fmt.Println(cardFloatalt)

	// Array and slice
	// Both are arrays , slice is ann array that can grow or shrink
	cards := []string{"Aces of Diamonds", newCard()}
	cards = append(cards, "Six of spades")
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Aces of spades"
}
