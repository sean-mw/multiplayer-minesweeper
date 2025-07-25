package game

import (
	"math/rand"
)

type Cell struct {
	IsMine     bool `json:"isMine"`
	IsRevealed bool `json:"isRevealed"`
	IsFlagged  bool `json:"isFlagged"`
	Adjacent   int  `json:"adjacent"`
}

type Board struct {
	Width    int      `json:"width"`
	Height   int      `json:"height"`
	Cells    [][]Cell `json:"cells"`
	Mines    int      `json:"mines"`
	Revealed int      `json:"revealed"`
}

var directions = []struct{ dx, dy int }{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func NewBoard(width, height int) *Board {
	cells := make([][]Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]Cell, width)
	}
	return &Board{
		Width:    width,
		Height:   height,
		Cells:    cells,
		Mines:    0,
		Revealed: 0,
	}
}

func (b *Board) PlaceMines(numMines int) {
	b.Mines = numMines
	placed := 0
	for placed < numMines {
		x := rand.Intn(b.Width)
		y := rand.Intn(b.Height)
		if !b.Cells[y][x].IsMine {
			b.Cells[y][x].IsMine = true
			placed++
		}
	}
}

func (b *Board) CalculateAdjacent() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			count := 0
			for _, d := range directions {
				nx, ny := x+d.dx, y+d.dy
				if nx >= 0 && nx < b.Width && ny >= 0 && ny < b.Height {
					if b.Cells[ny][nx].IsMine {
						count++
					}
				}
			}
			b.Cells[y][x].Adjacent = count
		}
	}
}

func (b *Board) Reveal(x, y int) {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return
	}
	cell := &b.Cells[y][x]
	if cell.IsRevealed || cell.IsFlagged {
		return
	}

	cell.IsRevealed = true
	b.Revealed += 1

	if cell.Adjacent == 0 && !cell.IsMine {
		for _, d := range directions {
			b.Reveal(x+d.dx, y+d.dy)
		}
	}
}

func (b *Board) ToggleFlag(x, y int) {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return
	}
	cell := &b.Cells[y][x]
	if cell.IsRevealed {
		return
	}
	cell.IsFlagged = !cell.IsFlagged
}
