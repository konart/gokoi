package main

import (
	"fmt"
)

func main() {

	pl1 := Player{Name: "Artem"}
	pl2 := Player{Name: "Alex"}
	newRoom := Room{playerOne: pl1, playerTwo: pl2}
	fmt.Printf("Rooom %v\n", newRoom)
	localDeck := new(Deck)
	localDeck.init()
	fmt.Printf("Deck %v\n", *localDeck)
	localDeck.Shuffle()
	fmt.Printf("Shuffled deck %v\n", *localDeck)
}
