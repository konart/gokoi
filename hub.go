package main

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

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			msg := []byte{'0'}
			client.send <- msg
			h.players[client] = true
		case client := <-h.unregister:
			if _, ok := h.players[client]; ok {
				delete(h.players, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.players {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.players, client)
				}
			}
		}
	}
}
