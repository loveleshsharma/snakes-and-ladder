package internal

import (
	"fmt"
)

type CLI struct {
}

func NewCLIReader() CLI {
	return CLI{}
}

func (c CLI) Read() ([]Snake, []Ladder, []Player, error) {
	var snakesCount, laddersCount, playersCount int

	snakes := make([]Snake, 0)
	ladders := make([]Ladder, 0)
	players := make([]Player, 0)

	fmt.Scanf("%d", &snakesCount)
	for i := 0; i < snakesCount; i++ {
		var head, tail int
		fmt.Scanf("%d %d", &head, &tail)
		snakes = append(snakes, NewSnake(head, tail))
	}

	fmt.Scanf("%d", &laddersCount)
	for i := 0; i < laddersCount; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)
		ladders = append(ladders, NewLadder(start, end))
	}

	fmt.Scanf("%d", &playersCount)
	for i := 0; i < playersCount; i++ {
		var id int
		var name string
		fmt.Scanf("%d %s", &id, &name)
		players = append(players, NewPlayer(id, name))
	}

	return snakes, ladders, players, nil
}
