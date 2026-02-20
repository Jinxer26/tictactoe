package game

import (
	"errors"
)

// File to handle the game board

type Cell rune

const (
	Empty Cell = ' '
	X     Cell = 'X'
	O     Cell = 'O'
)

type Board struct{
	Cells [3][3]Cell
}

func NewBoard() *Board{
	return &Board{}
}

func (b *Board) Place(row int, col int, c Cell ) error {
	
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("Out of Bounds, Specify valid column and row")
	}

	if b.Cells[row][col] != Empty {
		return errors.New("Cell occupied")
	}

	b.Cells[row][col] = c

}


