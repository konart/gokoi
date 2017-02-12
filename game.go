package main

type Table []Card

type Game struct {
	table Table
	players []*Player
	id int
	deck Deck
}
