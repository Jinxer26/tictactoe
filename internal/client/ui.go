package client

import (
	"fmt"

	"github.com/Jinxer26/tictactoe/internal/game"
)

func Render(g *game.Game) {

	fmt.Print("\033[2J\033[H") // clear screen

	fmt.Println("   0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d  ", i)
		for j := 0; j < 3; j++ {
			c := g.Board.Cells[i][j]
			if c == game.Empty {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", c)
			}
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  ---+---+---")
		}
	}
	fmt.Printf("\nTurn: %c\n", g.Turn)
}
