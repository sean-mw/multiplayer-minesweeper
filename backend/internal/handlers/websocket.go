package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: restrict origins in prod
	},
}

func RoomsWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) != 4 || parts[3] != "ws" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	roomId := parts[2]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	registerClient(roomId, conn)
	defer unregisterClient(roomId, conn)

	log.Printf("Client connected to room: %s", roomId)

	sendGameState(roomId, conn)

	for {
		var msg ClientMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		log.Printf("Received from client in room %s: %+v", roomId, msg)

		processClientAction(roomId, msg)
		broadcastGameState(roomId)
	}

	log.Printf("Client disconnected from room: %s", roomId)
}
