package main

func main() {

	// cards := deck{"Aces of Diamonds", newCard()}
	// cards = append(cards, "Six of spades")
	// cards := newDeck()
	// fmt.Println(cards)
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println(":::::::::::::hand::::::::::::::::::::::")
	// hand.print()
	// fmt.Println(":::::::::::::hand::::::::::::::::::::::")
	// fmt.Println(":::::::::::::remainingCards::::::::::::::::::::::")
	// remainingCards.print()
	// fmt.Println(":::::::::::::remainingCards::::::::::::::::::::::")
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffile()
	cards.print()
}

func newCard() string {
	return "Aces of spades"
}
