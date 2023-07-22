package main

import (
	"fmt"
	"time"
)

type Game struct {
	board   Board
	players []Player
	dice    Dice
}

func NewGame(board Board, dice Dice) Game {
	return Game{
		board:   board,
		players: make([]Player, 0),
		dice:    dice,
	}
}

func (g *Game) StartGame() error {
	fmt.Println("Starting the Game!")

	err := g.takeInput()
	if err != nil {
		return err
	}

	g.play()

	return nil
}

func (g *Game) play() {

	for i := 0; ; i++ {
		if i == len(g.players) {
			i = 0
		}
		player := &g.players[i]

		diceNumber := g.dice.roll()

		toPosition := player.getPosition() + diceNumber

		//if snake exists
		if snake, ok := g.board.IsSnakeExists(toPosition); ok {
			toPosition = g.getSnakeTailPos(snake)
			fmt.Println("snake bite!")
		}

		//if ladder exists
		if ladder, ok := g.board.IsLadderExists(toPosition); ok {
			toPosition = g.getLadderEndPos(ladder)
			fmt.Println("ladder climb!")
		}

		if toPosition <= 100 {
			fmt.Printf("%s rolled a %d and moved from %d to %d\n", player.name, diceNumber, player.posOnBoard, toPosition)
			g.board.MovePlayer(player, toPosition)
		} else {
			toPosition = player.getPosition()
			fmt.Printf("%s rolled a %d and moved from %d to %d\n", player.name, diceNumber, player.posOnBoard, toPosition)
		}

		if toPosition == 100 {
			fmt.Printf("%s wins the game", player.name)
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (g *Game) getSnakeTailPos(snake Snake) int {
	anotherSnake, exists := g.board.IsSnakeExists(snake.tailPosition)

	if !exists {
		return snake.tailPosition
	}

	return g.getSnakeTailPos(anotherSnake)
}

func (g *Game) getLadderEndPos(ladder Ladder) int {
	anotherLadder, exists := g.board.IsLadderExists(ladder.endPosition)

	if !exists {
		return ladder.endPosition
	}

	return g.getLadderEndPos(anotherLadder)
}

func (g *Game) takeInput() error {
	var snakes, ladders, players int

	fmt.Scanf("%d", &snakes)
	for i := 0; i < snakes; i++ {
		var head, tail int
		fmt.Scanf("%d %d", &head, &tail)
		fmt.Println(head, tail)

		newSnake := NewSnake(head, tail)
		if snakeErr := g.board.PutSnake(newSnake); snakeErr != nil {
			fmt.Println(snakeErr)
			return snakeErr
		}
	}

	fmt.Scanf("%d", &ladders)
	for i := 0; i < ladders; i++ {
		var start, end int
		fmt.Scanf("%d %d", &start, &end)

		newLadder := NewLadder(start, end)
		if ladderErr := g.board.PutLadder(newLadder); ladderErr != nil {
			fmt.Println(ladderErr)
			return ladderErr
		}
	}

	fmt.Scanf("%d", &players)
	for i := 0; i < players; i++ {
		var id int
		var name string
		fmt.Scanf("%d %s", &id, &name)

		newPlayer := NewPlayer(id, name)
		g.players = append(g.players, newPlayer)
	}
	return nil
}
