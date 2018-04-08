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

type card struct {
	Suit  string `json:"suit"`
	Group string `json:"group"`
}

// Stack of card pointers.
type deck []*card

// Adds a pointer to a card to the stack.
func (d *deck) add(c *card) {
	*d = append(*d, c)
}

// Pops an element from a stack.
func (d *deck) PickCard() *card {
	topCardIndex := len(*d) - 1
	c := (*d)[topCardIndex]
	*d = (*d)[:topCardIndex]
	return c
}

// Peek at the top element of the stack.
func (d deck) OpenCard() *card {
	topCardIndex := len(d) - 1
	c := d[topCardIndex]
	return c
}

// Shuffles the deck
func (d *deck) Shuffle() {
	tempDeck := make([]*card, 48)
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(48)
	for i, v := range list {
		tempDeck[i] = (*d)[v]
	}
	*d = tempDeck
}