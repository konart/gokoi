package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTemplate = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func main() {

	flag.Parse()
	hub := newHub()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	// New game (restrict to POST?)
	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		handleGames(hub, -1, w, r)
	})

	// Check for existing game (restrict to GET?)
	r.HandleFunc("/games/{gameId:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		gameId, err := strconv.ParseInt(mux.Vars(r)["gameId"], 10, 64)
		if err != nil {
			log.Fatal("Games Request Handler: ", err)
			return
		}
		handleGames(hub, gameId, w, r)
	})

	//r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(hub, -1, w, r)
	//})

	// Handle ws for a game
	r.HandleFunc("/ws/{gameId:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		gameId, err := strconv.ParseInt(mux.Vars(r)["gameId"], 10, 64)
		if err != nil {
			log.Fatal("Request Handler: ", err)
			return
		}
		serveWs(hub, gameId, w, r)
	})
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:15000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	//http.HandleFunc("/", serveHome)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(hub, w, r)
	//})
	//err := http.ListenAndServe(*addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	var game Game
	var deck Deck
	deck = make([]*Card, 48)
	deck.Prepare()
	for _, card := range deck {
		fmt.Println(*card)
	}

	game.deck = deck
}
