package logic

import (
	"testing"
	"fmt"
)

func TestCardsMap(t *testing.T) {
	cardTest := struct {
		suit string
		group string
		card string
	}{
		"Pine",
		"plain1",
		fmt.Sprintf(`{"suit":"%s","group":"%s"}`, "Pine", "plain1"),
	}
	card := &card{cardTest.suit, cardTest.group}

	cm := NewCardsMap()
	cm.add(card)
	got := cm.getCard(cardTest.card)
	if got.Group != card.Group {
		t.Errorf("Expected Group to be %s, got %s instead", card.Group, got.Group)
	}
	if got.Suit != card.Suit {
		t.Errorf("Expected Group to be %s, got %s instead", card.Suit, got.Suit)
	}
}
