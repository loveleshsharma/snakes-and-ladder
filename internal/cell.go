package internal

import (
	"errors"
)

type Cell struct {
	id      int
	snakes  []Snake
	ladders []Ladder
	players map[int]Player

	isEmpty bool
}

func NewCell(id int) Cell {
	return Cell{
		id:      id,
		snakes:  make([]Snake, 0),
		ladders: make([]Ladder, 0),
		players: make(map[int]Player),
		isEmpty: true,
	}
}

func (c *Cell) isCellEmpty() bool {
	return c.isEmpty
}

func (c *Cell) GetSnakes() []Snake {
	return c.snakes
}

func (c *Cell) GetLadders() []Ladder {
	return c.ladders
}

func (c *Cell) PutSnake(snake Snake) error {
	if len(c.ladders) > 0 {
		return errors.New("ladder already exists on this cell")
	}

	if len(c.snakes) == 0 {
		c.snakes = append(c.snakes, snake)
		return nil
	} else if len(c.snakes) < 2 {
		//check whether it's a tail
		if c.snakes[0].isTailPos(snake.headPosition) {
			c.snakes = append(c.snakes, snake)
		}
		return nil
	}

	return errors.New("snakes already exists on this cell")
}

func (c *Cell) PutLadder(ladder Ladder) error {
	if len(c.snakes) > 0 {
		return errors.New("snake already exists on this cell")
	}

	if len(c.ladders) == 0 {
		c.ladders = append(c.ladders, ladder)
		return nil
	} else if len(c.ladders) < 2 {
		if c.ladders[0].isStartPosition(ladder.endPosition) {
			c.ladders = append(c.ladders, ladder)
		}
		return nil
	}

	return errors.New("ladder already exists on this cell")
}

func (c *Cell) PutPlayer(player *Player) {
	if c.players == nil {
		c.players = make(map[int]Player)
	}

	c.players[player.id] = *player
}

func (c *Cell) RemovePlayer(player Player) {
	delete(c.players, player.id)
}
