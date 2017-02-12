package main

import (
	"fmt"
)

func main() {

	var game Game
	var deck Deck
	deck = make([]*Card, 48)
	deck.Prepare()
	for _, card := range deck {
		fmt.Println(*card)
	}

	game.deck = deck
}
