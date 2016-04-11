package main

import (
	"time"
	"math/rand"
)

type Card struct {
	Suit  int
	Value int
}

type Deck struct {
	Cards [48]Card
}

func (d *Deck) Shuffle() {
	var tempDeck [len(d.Cards)]Card
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(len(d.Cards))
	for i, v := range list{
		tempDeck[i] = d.Cards[v]
	}
	d.Cards = tempDeck
}