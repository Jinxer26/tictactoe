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
			fmt.Println("Invalid Input")
			continue
		}

		if err := g.Move(r,c); err != nil {
			fmt.Println("Error: ", err)
			continue
		}

	}

	Render(g)

	fmt.Printf("\n Winner : %c \n ",g.Winner)
}
