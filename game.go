package main

import "github.com/konart/gokoi/logic"

func StartTheGame() {
	deck := logic.NewDeck()
	cardsMap := logic.NewCardsMap()
	logic.InitDeckAndMap(deck, cardsMap)
}
