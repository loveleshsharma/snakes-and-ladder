package internal

type Snake struct {
	headPosition int
	tailPosition int
}

func NewSnake(headPos, tailPos int) Snake {
	return Snake{
		headPosition: headPos,
		tailPosition: tailPos,
	}
}

func (s Snake) isHeadPos(position int) bool {
	return position == s.headPosition
}

func (s Snake) isTailPos(position int) bool {
	return position == s.tailPosition
}

func (s Snake) GetType() UnitType {
	return SNAKE
}
