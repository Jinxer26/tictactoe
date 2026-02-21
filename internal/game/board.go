package game

import (
	"errors"
	"github.com/Jinxer26/tictactoe/internal/logger"
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
	  b := &Board{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            b.Cells[i][j] = Empty
        }
    }
    return b
}

func (b *Board) Place(row int, col int, c Cell ) error {
	
	if row < 0 || row > 2 || col < 0 || col > 2 {
		logger.Log.Error("Out of Bounds", "row",row,"col",col)
		return errors.New("Out of Bounds")
	}

	if b.Cells[row][col] != Empty {
		logger.Log.Error("Cell Occupied", "value", string(b.Cells[row][col]))
		return errors.New("Cell occupied")
	}

	b.Cells[row][col] = c
  
	return nil
}


