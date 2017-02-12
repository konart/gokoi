package main

type Table []*Card

type Game struct {
	table Table
	players []*Player // we need this?
	id int
	deck Deck
}
