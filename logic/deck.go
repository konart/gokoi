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

type deck struct {
	cards []*card
}

func (d *deck) add(c *card) {
	d.cards = append(d.cards, c)
}

func (d *deck) PickCard() *card {
	topCardIndex := len(d.cards) - 1
	c := d.cards[topCardIndex]
	d.cards = d.cards[:topCardIndex]
	return c
}

func (d deck) OpenCard() *card {
	topCardIndex := len(d.cards) - 1
	c := d.cards[topCardIndex]
	return c
}

func (d *deck) Shuffle() {
	tempDeck := make([]*card, 48)
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(48)
	for i, v := range list {
		tempDeck[i] = (d.cards)[v]
	}
	d.cards = tempDeck
}

func (d deck) hasCard(c *card) bool {
	return d.OpenCard() == c
}

func initDeckAndHands(d *deck, cm *CardsMap, hand1, hand2 *hand) {
	for key, value := range cards {
		for _, group := range value {
			card := &card{key, group}
			cm.add(card)
			d.add(card)
			hand1.add(card)
			hand2.add(card)
		}
	}
}

func PrepareDeck(d *deck) {
	d.Shuffle()
}