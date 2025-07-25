package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/sean-mw/multiplayer-minesweeper/backend/internal/game"
)

var (
	gamesMu sync.Mutex
	games   = make(map[string]*game.Game)
)

func RoomsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		log.Println("Creating a new room")
		roomId := uuid.NewString()
		gamesMu.Lock()
		games[roomId] = game.NewGame(roomId, 10, 10, 10)
		gamesMu.Unlock()

		resp := map[string]string{"roomId": roomId}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
