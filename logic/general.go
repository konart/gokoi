package logic

import (
	"encoding/json"
	"log"
)

// Initiates global deck and card map.
func InitDeckAndMap(d deck, cm CardsMap) {
	for key, value := range cards {
		for _, group := range value {
			card := &card{key, group}
			cm.add(card)
			d.add(card)
		}
	}
}

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

// A simple map to keep JSON representation of cards
type CardsMap map[string]*card

// Returns a new CardsMap struct.
func NewCardsMap() CardsMap {
	m := make(map[string]*card)
	return m
}

// Adds new entry to a CardsMap
// Key will be a JSON representation of a card struct.
func (c CardsMap) add(card *card) {
	cardString, err := json.Marshal(card)
	if err != nil {
		log.Panic(err)
	}
	c[string(cardString)] = card
}

// Returns a pointer to a card for a given JSON string
func (c CardsMap) getCard(s string) *card {
	return c[s]
}