package main

// "fmt"

// "go_crud/greeting"
// "go_crud/hello"

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// ===============================

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// ===============================

	// cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()

	// ===============================

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
