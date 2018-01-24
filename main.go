package main

func main() {
	//var card = "Ace of Spades"
	//card := newCard()
	//slice
	cards := newDeck()
	hand, ramainigCards := deal(cards, 5)

	cards.print()
	hand.print()
	ramainigCards.print()

	cards.saveToFile("./dump")
}
