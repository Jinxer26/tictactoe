package main

import(
	"github.com/Jinxer26/tictactoe/internal/client"
	"github.com/Jinxer26/tictactoe/internal/logger"
)

func main(){
	logger.Init()
	client.RunLocalGame()
}
