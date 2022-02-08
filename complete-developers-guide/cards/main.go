package main

func main() {
	// Create a new deckObject and shuffle it
	d := newDeck()
	d.shuffle()

	// Give a player hand 5 cards from the shuffled deck and print it
	playerHand, d := deal(d, 5)
	playerHand.print()
}
