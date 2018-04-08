package logic

import (
	"math/rand"
	"time"
)

var cards = map[string][]string{
	"Pine":          {"plain1", "plain2", "ribbon", "bright"},
	"Plum":          {"plain1", "plain2", "ribbon", "animal"},
	"Sakura":        {"plain1", "plain2", "ribbon", "bright"},
	"Wisteria":      {"plain1", "plain2", "ribbon", "animal"},
	"Iris":          {"plain1", "plain2", "ribbon", "animal"},
	"Peony":         {"plain1", "plain2", "ribbon", "animal"},
	"Clover":        {"plain1", "plain2", "ribbon", "animal"},
	"Pampas":        {"plain1", "plain2", "animal", "bright"},
	"Chrysanthemum": {"plain1", "plain2", "ribbon", "animal"},
	"Maple":         {"plain1", "plain2", "ribbon", "animal"},
	"Willow":        {"plain1", "ribbon", "animal", "bright"},
	"Paulownia":     {"plain1", "plain2", "plain3", "bright"},
}

type Card struct {
	Suit  string `json:"suit"`
	Group string `json:"group"`
}

// Stack of Card pointers.
type Deck []*Card

func NewDeck() Deck {
	return Deck{}
}

// Adds a pointer to a Card to the stack.
func (d *Deck) add(c *Card) {
	*d = append(*d, c)
}

// Pops an element from a stack.
func (d *Deck) PickCard() *Card {
	topCardIndex := len(*d) - 1
	c := (*d)[topCardIndex]
	*d = (*d)[:topCardIndex]
	return c
}

// Peek at the top element of the stack.
func (d Deck) OpenCard() *Card {
	topCardIndex := len(d) - 1
	c := d[topCardIndex]
	return c
}

// Shuffles the Deck
func (d *Deck) Shuffle() {
	tempDeck := make([]*Card, 48)
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(48)
	for i, v := range list {
		tempDeck[i] = (*d)[v]
	}
	*d = tempDeck
}