package game

import (
	"sync"
)

type GameStatus int

const (
	Pending GameStatus = iota
	Won
	Lost
)

type Game struct {
	sync.Mutex

	ID     string
	Board  *Board
	Status GameStatus
}

func NewGame(id string, width, height, mines int) *Game {
	board := NewBoard(width, height)
	board.PlaceMines(mines)
	board.CalculateAdjacent()

	return &Game{
		ID:     id,
		Board:  board,
		Status: Pending,
	}
}

func (g *Game) RevealCell(x, y int) {
	g.Lock()
	defer g.Unlock()

	if g.Status != Pending {
		return
	}

	cell := &g.Board.Cells[y][x]
	if cell.IsRevealed || cell.IsFlagged {
		return
	}

	g.Board.Reveal(x, y)

	if cell.IsMine {
		g.Status = Lost
		return
	}

	if g.Board.Revealed == (g.Board.Width*g.Board.Height - g.Board.Mines) {
		g.Status = Won
	}
}

func (g *Game) ToggleFlag(x, y int) {
	g.Lock()
	defer g.Unlock()

	if g.Status != Pending {
		return
	}

	g.Board.ToggleFlag(x, y)
}
