package handlers

import (
	"log"
)

func processClientAction(roomId string, msg ClientMessage) {
	gameInstance := getGame(roomId)
	if gameInstance == nil {
		log.Printf("Game not found for room %s", roomId)
		return
	}

	switch msg.Action {
	case "reveal":
		gameInstance.RevealCell(msg.X, msg.Y)
	case "flag":
		gameInstance.ToggleFlag(msg.X, msg.Y)
	default:
		log.Printf("Unknown action: %s", msg.Action)
	}
}
