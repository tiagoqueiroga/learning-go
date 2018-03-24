package main

func main() {
	// Slice of cards(Can be changed)
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
