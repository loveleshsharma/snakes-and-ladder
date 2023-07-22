package main

import "fmt"

func main() {
	game := NewGame(NewBoard(), NewDice())

	if err := game.StartGame(); err != nil {
		fmt.Println(err)
	}
}
