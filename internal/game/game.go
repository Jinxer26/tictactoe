package game

import (
	"errors"
	"github.com/Jinxer26/tictactoe/internal/logger"
)

type Game struct {
    Board  *Board
    Turn   Cell
    Over   bool
    Winner Cell
}

func NewGame() *Game {
	  logger.Log.Info("Setting up new game")
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
		logger.Log.Info("Move Played","player",string(g.Turn),"row",row,"col",col)

    w := g.Board.Winner()
    if w != Empty {
        g.Over = true
        g.Winner = w
				logger.Log.Info("Game Over ! ", "winner" , string(g.Turn))
        return nil
    }

		d := g.Board.Draw()
		if d{
			g.Over = true
			g.Winner = Empty
			logger.Log.Info("Game Draw, nobody wins")
			return errors.New("Draw")
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


