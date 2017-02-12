package main

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	players map[*Player]bool
	// Active games
	games map[*Game]bool

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
		games:      make(map[*Game]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.players[client] = true
			// here we should create a new game instance,
			// add it to the active games pool
			// tie player to the game
			game := &Game{
				table:   make([]*Card, 0),
				players: make([]*Player, 2),
				id:      len(h.games), // should be unique in future
				deck:    make([]*Card, 48),
			}
			h.games[game] = true

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
