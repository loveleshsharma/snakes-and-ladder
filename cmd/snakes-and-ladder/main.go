package main

import (
	"fmt"
	"github.com/loveleshsharma/snakesandladder/internal"
)

func main() {
	game := internal.NewGame(internal.NewBoard(), internal.NewDice(), internal.NewFileReader("./test-files/file2.txt"))

	if err := game.StartGame(); err != nil {
		fmt.Println(err)
	}
}
