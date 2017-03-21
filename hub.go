package main

import (
	"net/http"
	"log"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	players map[*Player]bool

	// Active games
	games map[int64]*Game

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Player

	// Unregister requests from clients.
	unregister chan *Player
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		players:    make(map[*Player]bool),
		games:      make(map[int64]*Game),
	}
}

func (h *Hub) gameExists(id int64) bool {
	game := h.games[id]
	if game != nil {
		return true
	} else {
		return false
	}
}

func (h *Hub) findGame(id int64) (game *Game) { // old (game *Game, err error)
	game = h.games[id]
	if game == nil {
		game = &Game{}
	}
	return game //, err
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, gameId int64, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Find game
	game := hub.findGame(gameId)

	// Create client
	player := &Player{game: game, conn: conn, send: make(chan []byte, 256)}

	// Add to game
	player.game.register <- player

	go player.writePump()
	player.readPump()

	log.Printf("GameID: %d", game.id)
}