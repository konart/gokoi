package main

import (
	"math/rand"
	"time"
)

type Card struct {
	Suit  int
	Value int
}

type CardArray struct {
	Cards [48]Card
}

type Combination struct {
	// TODO: this one needs something to tie it to cards
	cardsLeft int
}

type Deck CardArray

type Hand CardArray

type OpenHand CardArray

func (d *Deck) Shuffle() {
	var tempDeck [len(d.Cards)]Card
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(len(d.Cards))
	for i, v := range list {
		tempDeck[i] = d.Cards[v]
	}
	d.Cards = tempDeck
}

func (d *Deck) init() {
	i := 0
	for j := 0; j < 12; j++ {
		for k := 0; k < 4; k++ {
			playingCard := new(Card)
			playingCard.Value = k
			playingCard.Suit = j
			d.Cards[i] = *playingCard
			i++
		}
	}
}

func (o *OpenHand) link(*Card, *Combination) {

}
