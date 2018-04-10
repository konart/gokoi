package main

import "github.com/konart/gokoi/logic"

func initTheGame() {
	deck := logic.NewDeck()
	cardsMap := logic.NewCardsMap()
	logic.InitDeckAndMap(deck, cardsMap)
}

// Creates a game hub.
// All communications between the two players will be here
func createGameHub() {

}

func addPlayerToHub(p *logic.Player, h *logic.Hub) {

}