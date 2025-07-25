package handlers

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sean-mw/multiplayer-minesweeper/backend/internal/game"
)

var (
	clientsMu sync.Mutex
	clients   = make(map[string]map[*websocket.Conn]bool)
)

func registerClient(roomId string, conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	if clients[roomId] == nil {
		clients[roomId] = make(map[*websocket.Conn]bool)
	}
	clients[roomId][conn] = true
}

func unregisterClient(roomId string, conn *websocket.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	if conns, ok := clients[roomId]; ok {
		delete(conns, conn)
		if len(conns) == 0 {
			delete(clients, roomId)
		}
	}
	conn.Close()
}

func sendGameState(roomId string, conn *websocket.Conn) {
	gameInstance := getGame(roomId)
	if gameInstance == nil {
		return
	}

	resp := ServerMessage{
		Board:  gameInstance.Board,
		Status: gameInstance.Status,
	}

	err := conn.WriteJSON(resp)
	if err != nil {
		log.Println("Failed to send game state:", err)
	}
}

func broadcastGameState(roomId string) {
	clientsMu.Lock()
	conns, ok := clients[roomId]
	clientsMu.Unlock()

	if !ok {
		return
	}

	gameInstance := getGame(roomId)
	if gameInstance == nil {
		return
	}

	resp := ServerMessage{
		Board:  gameInstance.Board,
		Status: gameInstance.Status,
	}

	for conn := range conns {
		if err := conn.WriteJSON(resp); err != nil {
			log.Println("Error broadcasting to client:", err)
		}
	}
}

func getGame(roomId string) *game.Game {
	gamesMu.Lock()
	defer gamesMu.Unlock()
	return games[roomId]
}
