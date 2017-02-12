package main

import (
	"fmt"
)

func main() {

	var game Game
	var deck Deck

	deck = make([]*Card, 48)
	i := 0
	for key, value := range cards {
		for _, group := range value {
			deck[i] = &Card{key, group}
			i++
		}
	}
	deck.Shuffle()
	for _, card := range deck {
		fmt.Println(*card)
	}

	_ = *deck.PickCard()
	_ = *deck.PickCard()
	_ = *deck.PickCard()

	fmt.Print(len(deck))



	game.deck = deck
}
