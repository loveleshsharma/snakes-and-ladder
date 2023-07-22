package main

import (
	"errors"
)

type Board struct {
	cellGrid [][]Cell
}

func NewBoard() Board {
	cellGrid := [][]Cell{
		{NewCell(100), NewCell(99), NewCell(98), NewCell(97), NewCell(96), NewCell(95), NewCell(94), NewCell(93), NewCell(92), NewCell(91)},
		{NewCell(81), NewCell(82), NewCell(83), NewCell(84), NewCell(85), NewCell(86), NewCell(87), NewCell(88), NewCell(89), NewCell(90)},
		{NewCell(80), NewCell(79), NewCell(78), NewCell(77), NewCell(76), NewCell(75), NewCell(74), NewCell(73), NewCell(72), NewCell(71)},
		{NewCell(61), NewCell(62), NewCell(63), NewCell(64), NewCell(65), NewCell(66), NewCell(67), NewCell(68), NewCell(69), NewCell(70)},
		{NewCell(60), NewCell(59), NewCell(58), NewCell(57), NewCell(56), NewCell(55), NewCell(54), NewCell(53), NewCell(52), NewCell(51)},
		{NewCell(41), NewCell(42), NewCell(43), NewCell(44), NewCell(45), NewCell(46), NewCell(47), NewCell(48), NewCell(49), NewCell(50)},
		{NewCell(40), NewCell(39), NewCell(38), NewCell(37), NewCell(36), NewCell(35), NewCell(34), NewCell(33), NewCell(32), NewCell(31)},
		{NewCell(21), NewCell(22), NewCell(23), NewCell(24), NewCell(25), NewCell(26), NewCell(27), NewCell(28), NewCell(29), NewCell(30)},
		{NewCell(20), NewCell(19), NewCell(18), NewCell(17), NewCell(16), NewCell(15), NewCell(14), NewCell(13), NewCell(12), NewCell(11)},
		{NewCell(1), NewCell(2), NewCell(3), NewCell(4), NewCell(5), NewCell(6), NewCell(7), NewCell(8), NewCell(9), NewCell(10)},
	}

	return Board{
		cellGrid: cellGrid,
	}
}

func (b *Board) IsSnakeExists(position int) (Snake, bool) {
	x, y := b.FindCoordinatesInGrid(position)

	snakes := b.cellGrid[x][y].GetSnakes()

	for _, snake := range snakes {
		if snake.isHeadPos(position) {
			return snake, true
		}
	}

	return Snake{}, false
}

func (b *Board) IsLadderExists(position int) (Ladder, bool) {
	x, y := b.FindCoordinatesInGrid(position)

	ladders := b.cellGrid[x][y].GetLadders()

	for _, ladder := range ladders {
		if ladder.isStartPosition(position) {
			return ladder, true
		}
	}

	return Ladder{}, false
}
func (b *Board) PutSnake(snake Snake) error {
	if snake.headPosition < snake.tailPosition {
		return errors.New("invalid snake position")
	}

	headX, headY := b.FindCoordinatesInGrid(snake.headPosition)
	tailX, tailY := b.FindCoordinatesInGrid(snake.tailPosition)

	if err := b.cellGrid[headX][headY].PutSnake(snake); err != nil {
		return err
	}

	if err := b.cellGrid[tailX][tailY].PutSnake(snake); err != nil {
		return err
	}

	return nil
}

func (b *Board) PutLadder(ladder Ladder) error {
	if ladder.startPosition > ladder.endPosition {
		return errors.New("invalid ladder position")
	}

	startX, startY := b.FindCoordinatesInGrid(ladder.startPosition)
	endX, endY := b.FindCoordinatesInGrid(ladder.endPosition)

	//if b.cellGrid[startX][startY].unit != nil || b.cellGrid[endX][endY].unit != nil {
	//	return errors.New("clash, other unit already exists at this position")
	//}
	//
	//b.cellGrid[startX][startY].unit = ladder
	//b.cellGrid[endX][endY].unit = ladder

	if err := b.cellGrid[startX][startY].PutLadder(ladder); err != nil {
		return err
	}

	if err := b.cellGrid[endX][endY].PutLadder(ladder); err != nil {
		return err
	}

	return nil
}

func (b *Board) MovePlayer(player *Player, toPosition int) {
	oldX, oldY := b.FindCoordinatesInGrid(player.posOnBoard)
	newX, newY := b.FindCoordinatesInGrid(toPosition)

	b.cellGrid[oldX][oldY].RemovePlayer(*player)
	b.cellGrid[newX][newY].PutPlayer(player)

	player.setPosition(toPosition)
}

/*
<10 - row 9 inc
<20 - row 8 dec
<30 - row 7 inc
<40 - row 6 dec
<50 - row 5 inc
<60 - row 4 dec
<70 - row 3 inc
<80 - row 2 dec
<90 - row 1 inc
<100 - row 0 dec
*/
func (b *Board) FindCoordinatesInGrid(pos int) (int, int) {
	var x, y int

	switch {
	case pos <= 10:
		mod := pos % 10
		y = 9
		x = getXCoordinateForEvenRow(mod)
	case pos <= 20:
		mod := pos % 10
		y = 8
		x = getXCoordinateForOddRow(mod)
	case pos <= 30:
		mod := pos % 10
		y = 7
		x = getXCoordinateForEvenRow(mod)
	case pos <= 40:
		mod := pos % 10
		y = 6
		x = getXCoordinateForOddRow(mod)
	case pos <= 50:
		mod := pos % 10
		y = 5
		x = getXCoordinateForEvenRow(mod)
	case pos <= 60:
		mod := pos % 10
		y = 4
		x = getXCoordinateForOddRow(mod)
	case pos <= 70:
		mod := pos % 10
		y = 3
		x = getXCoordinateForEvenRow(mod)
	case pos <= 80:
		mod := pos % 10
		y = 2
		x = getXCoordinateForOddRow(mod)
	case pos <= 90:
		mod := pos % 10
		y = 1
		x = getXCoordinateForEvenRow(mod)
	case pos <= 100:
		mod := pos % 10
		y = 0
		x = getXCoordinateForOddRow(mod)
	}

	return x, y
}

func getXCoordinateForOddRow(mod int) int {
	var x int
	if mod == 0 {
		x = 0
	} else {
		x = 10 - mod
	}

	return x
}

func getXCoordinateForEvenRow(mod int) int {
	var x int
	if mod == 0 {
		x = 9
	} else {
		x = mod - 1
	}

	return x
}
