package main

import "fmt"

func main() {
	cards := newDeck()

	newDeck, remainDeck := deal(cards, 5)

	newDeck.print()
	fmt.Println("------")
	remainDeck.print()
}
