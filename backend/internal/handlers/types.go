package handlers

import "github.com/sean-mw/multiplayer-minesweeper/backend/internal/game"

type ClientMessage struct {
	Action string `json:"action"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

type ServerMessage struct {
	Board  *game.Board     `json:"board"`
	Status game.GameStatus `json:"status"`
}
