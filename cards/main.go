package main

// we can initialize the variable outside the function
// but we can not assign a value to that variable

// var desckSize int

func main() {
	// var card string = "Ace of Spades"

	// var card = "Ace of Spades"

	// card := "Ace of Spades"  // this is only valid for first declaration

	// card = "Joker of Spades" // second or more assign use '='

	// card := newCard()

	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("Hand:")
	// hand.print()
	// fmt.Println("remainingDeck:")
	// remainingDeck.print()

	// cards := newDeck()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")
	cards := newDeck()

	cards.shuffle()
	cards.print()
}
