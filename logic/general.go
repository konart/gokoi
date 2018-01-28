package logic

import (
	"encoding/json"
	"log"
)

type cardHolder struct {
	cards map[*card]bool
}

func (t *cardHolder) add(c *card) {
	t.cards[c] = true
}

func (t cardHolder) hasCard(c *card) bool {
	return t.cards[c]
}

type cardHolderInterface interface {
	hasCard(*card) bool
}

func DealCards(d deck, t *table, p1, p2 *Player) {
	for i := 0; i < 8; i++ {
		p1.closedHand.add(d.PickCard())
		p2.closedHand.add(d.PickCard())
		t.add(d.PickCard())
	}
}

func CheckCards(c *card, ch cardHolderInterface) bool {
	if !ch.hasCard(c) {
		return false
	}
	return true
}

type CardsMap struct {
	cards map[string]*card
}

func (c *CardsMap) add(card *card) {
	cardString, err := json.Marshal(card)
	if err != nil {
		log.Panic(err)
	}
	c.cards[string(cardString)] = card
}

func (c CardsMap) getCard(s string) *card {
	return c.cards[s]
}