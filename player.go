package main

import "github.com/gorilla/websocket"

type Hand []*Card

type Combination interface {
	Satisfies(c *Card) bool
}

type Yaku struct {
	combinesWith *Yaku
	blocks       *Yaku
	collected    bool
}

func (y *Yaku) Satisfies(c *Card) bool {
	return true
}

type Player struct {
	game *Game
	conn *websocket.Conn
	send chan []byte
	openHand  Hand
	closedHand Hand
	name string
	id int
}
