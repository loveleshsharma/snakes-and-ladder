package main

import (
	"fmt"
	"github.com/loveleshsharma/snakesandladder/internal"
)

func main() {
	game := internal.NewGame(internal.NewBoard(), internal.NewDice())

	if err := game.StartGame(); err != nil {
		fmt.Println(err)
	}
}
