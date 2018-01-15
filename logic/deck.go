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
	suit  string
	group string
}

type deck []*card

func (d *deck) PickCard() *card {
	topCardIndex := len(*d) - 1
	c := (*d)[topCardIndex]
	*d = (*d)[:topCardIndex]
	return c
}

func (d deck) OpenCard() *card {
	topCardIndex := len(d) - 1
	c := d[topCardIndex]
	return c
}

func (d *deck) Shuffle() {
	tempDeck := make([]*card, 48)
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(48)
	for i, v := range list {
		tempDeck[i] = (*d)[v]
	}
	*d = tempDeck
}

func (d deck) HasCard(c *card) bool {
	for i := range d {
		if d[i].group == c.group && d[i].suit == c.suit {
			return true
		}
	}
	return false
}

func createNewDeck() *deck {
	deck := deck{}
	i := 0
	for key, value := range cards {
		for _, group := range value {
			deck[i] = &card{key, group}
			i++
		}
	}
	return &deck
}

func PrepareDeck() *deck {
	deck := createNewDeck()
	deck.Shuffle()
	return deck
}