package temp

import (
	"net/http"
	"strconv"
	"fmt"
	"log"
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

func (g *Game) hasBothPlayers() bool {
	if len(g.players) == 2 {
		return true
	} else {
		return false
	}
}

func (g *Game) hasEmptySlots() bool {
	if len(g.players) < 2 {
		return true
	} else {
		return false
	}
}

func newGame(h *Hub) (id int64) {
	game := &Game{
		id:         1,
		deck:       make([]*Card, 48),
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
			game := hub.findGame(gameId)
			if game.hasEmptySlots() {
				w.Write([]byte(strconv.FormatInt(gameId, 16))) // replace with json
			} else {
				log.Printf("The game with id %d has all slots full", gameId)
				w.Write([]byte(fmt.Sprintf("The game with id %d has all slots full", gameId)))  // replace with json
			}
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
			g.players[player] = true
			message := fmt.Sprintf("Player %d connected", len(g.players))
			g.sendToBoth(message)
			if g.hasBothPlayers() {
				g.deck.Prepare()
				card := g.deck.OpenCard()
				message := fmt.Sprintf("{suit: %s, group: %s}", card.suit, card.group)
				g.sendToBoth(message)
			}
		case player := <-g.unregister:
			if _, ok := g.players[player]; ok {
				delete(g.players, player)
				close(player.send)
			}
		case message := <-g.broadcast:
			for player := range g.players {
				select {
				case player.send <- message:
				default:
					close(player.send)
					delete(g.players, player)
				}
			}
		}
	}
}

func (game *Game) sendToPlayer(player *Player, message string) {
	player.send <- []byte(message)
}

func (game *Game) sendToBoth(message string) {
	for player := range game.players {
		player.send <- []byte(message)
	}
}
