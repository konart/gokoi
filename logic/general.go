package logic

import (
	"encoding/json"
	"log"
)

// Initiates global Deck and Card map.
func InitDeckAndMap(d Deck, cm CardsMap) {
	for key, value := range cards {
		for _, group := range value {
			card := &Card{key, group}
			cm.add(card)
			d.add(card)
		}
	}
}

type cardHolder struct {
	cards map[*Card]bool
}

func (t *cardHolder) add(c *Card) {
	t.cards[c] = true
}

func (t cardHolder) hasCard(c *Card) bool {
	return t.cards[c]
}

type cardHolderInterface interface {
	hasCard(*Card) bool
}

func DealCards(d Deck, t *table, p1, p2 *Player) {
	for i := 0; i < 8; i++ {
		p1.closedHand.add(d.PickCard())
		p2.closedHand.add(d.PickCard())
		t.add(d.PickCard())
	}
}

func CheckCards(c *Card, ch cardHolderInterface) bool {
	if !ch.hasCard(c) {
		return false
	}
	return true
}

// A simple map to keep JSON representation of cards
type CardsMap map[string]*Card

// Returns a new CardsMap struct.
func NewCardsMap() CardsMap {
	m := make(map[string]*Card)
	return m
}

// Adds new entry to a CardsMap
// Key will be a JSON representation of a Card struct.
func (c CardsMap) add(card *Card) {
	cardString, err := json.Marshal(card)
	if err != nil {
		log.Panic(err)
	}
	c[string(cardString)] = card
}

// Returns a pointer to a Card for a given JSON string
func (c CardsMap) getCard(s string) *Card {
	return c[s]
}