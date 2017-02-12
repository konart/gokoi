package main

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
	suit  string
	group string
}

type Deck []*Card

func (d *Deck) Prepare() {

	//d = make([]*Card, 48)
	i := 0
	for key, value := range cards {
		for _, group := range value {
			(*d)[i] = &Card{key, group}
			i++
		}
	}
	d.Shuffle()
}

func (d *Deck) PickCard() (c *Card) {
	topCardIndex := len(*d) - 1
	c = (*d)[topCardIndex]
	*d = (*d)[:topCardIndex]
	return
}

func (d *Deck) OpenCard() (c *Card) {
	topCardIndex := len(*d) - 1
	c = (*d)[topCardIndex]
	return
}

func (d *Deck) Shuffle() {
	tempDeck := make([]*Card, 48)
	rand.Seed(time.Now().UTC().UnixNano())
	list := rand.Perm(48)
	for i, v := range list {
		tempDeck[i] = (*d)[v]
	}
	*d = tempDeck
}
