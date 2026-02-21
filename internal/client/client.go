package client

import (
	"github.com/Jinxer26/tictactoe/internal/game"
	"fmt"
)

func RunLocalGame() {
	g := game.NewGame()

	for !g.Over{

		Render(g)
		
		r, c, err := ReadMove()
		if err != nil {
			continue
		}

		if err := g.Move(r,c); err != nil {
			continue
		}

	}

	Render(g)

	fmt.Printf("\n Winner : %c \n ",g.Winner)
}
