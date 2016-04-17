package main

type Room struct {
	playerOne Player
	playerTwo Player
	deck      Deck
}

type Player struct {
	Name     string
	hand     Deck
	openHand Deck
}
