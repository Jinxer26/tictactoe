package game

import (
	"errors"
)

type Game struct {
    Board  *Board
    Turn   Cell
    Over   bool
    Winner Cell
}

func NewGame() *Game {
    return &Game{
        Board: NewBoard(),
        Turn:  X,
    }
}

func (g *Game) Move(row, col int) error {
    if g.Over {
        return errors.New("Game Over")
    }

    if err := g.Board.Place(row, col, g.Turn); err != nil {
        return err
    }

    w := g.Board.Winner()
    if w != Empty {
        g.Over = true
        g.Winner = w
        return nil
    }

    g.Turn = opposite(g.Turn)
    return nil
}

func opposite(c Cell) Cell{
	if c == X {
		return O
	} 
	return X
}


