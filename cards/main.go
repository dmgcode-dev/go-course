package main

func main() {
	//cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards.print()
	//hand, remainingCards := deal(cards, 3)

	//hand.print()
	//remainingCards.print()

}
