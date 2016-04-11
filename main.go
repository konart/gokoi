package main

import (
	"fmt"
)

func main() {
	localDeck := new(Deck)
	i := 0
	for j := 0; j < 12; j++ {
		for k := 0; k < 4; k++ {
			playingCard := new(Card)
			playingCard.Value = k
			playingCard.Suit = j
			localDeck.Cards[i] = *playingCard
			i++
		}
	}

	fmt.Printf("%v\n", *localDeck)
	localDeck.Shuffle()
	fmt.Printf("%v\n", *localDeck)
}
