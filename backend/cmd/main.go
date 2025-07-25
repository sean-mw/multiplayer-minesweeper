package main

import (
	"log"
	"net/http"

	"github.com/sean-mw/multiplayer-minesweeper/backend/internal/handlers"
)


func main() {
	http.HandleFunc("/rooms", handlers.RoomsHandler)
	http.HandleFunc("/rooms/", handlers.RoomsWebSocketHandler)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error listening and serving:", err)
	}
}
