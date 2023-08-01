package main

import (
	"fmt"
	"github.com/loveleshsharma/snakesandladder/internal"
)

/*
	PENDING
	1. input via FILE and CLI
	2. introduce OOPS concepts
*/

func main() {
	game := internal.NewGame(internal.NewBoard(), internal.NewDice(), internal.NewFileReader("./test-files/file1.txt"))

	if err := game.StartGame(); err != nil {
		fmt.Println(err)
	}
}
