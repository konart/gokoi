package main

import (
	"net/http"
	"strconv"
)

type Table []*Card

type Game struct {
	table Table
	id    int64
	deck  Deck

	// Registered clients.
	players map[*Player]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Player

	// Unregister requests from clients.
	unregister chan *Player
}

func newGame(h *Hub) (id int64) {
	game := &Game{
		id:         1,
		register:   make(chan *Player),
		unregister: make(chan *Player),
		broadcast:  make(chan []byte),
		players:    make(map[*Player]bool),
	} // generate uuid
	go game.start()
	h.games[game.id] = game
	id = game.id
	return id
}

func handleGames(hub *Hub, gameId int64, w http.ResponseWriter, r *http.Request) {
	if gameId >= 0 {
		// Find game if we have gameId
		if hub.gameExists(gameId) {
			w.Header().Set("Server", "A Go Web Server")
			w.WriteHeader(200)
			w.Write([]byte("Found: "))
			w.Write([]byte(strconv.FormatInt(gameId, 16))) // replace with json
		} else {
			w.Header().Set("Server", "A Go Web Server")
			w.WriteHeader(200)
			w.Write([]byte("false")) // replace with json
		}
	} else {
		// Create new game if not
		gameId := newGame(hub)
		w.Header().Set("Server", "A Go Web Server")
		w.WriteHeader(200)
		w.Write([]byte(strconv.FormatInt(gameId, 16)))
	}
}

func (g *Game) start() {
	for {
		select {
		case player := <-g.register:
			msg := []byte("registered")
			player.send <- msg
			g.players[player] = true
		case player := <-g.unregister:
			if _, ok := g.players[player]; ok {
				delete(g.players, player)
				close(player.send)
			}
		case message := <-g.broadcast:
			for client := range g.players {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(g.players, client)
				}
			}
		}
	}
}
