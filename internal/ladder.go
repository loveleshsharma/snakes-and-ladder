package internal

type Ladder struct {
	startPosition int
	endPosition   int
}

func NewLadder(startPos, endPos int) Ladder {
	return Ladder{
		startPosition: startPos,
		endPosition:   endPos,
	}
}

func (l Ladder) isStartPosition(position int) bool {
	return l.startPosition == position
}

func (l Ladder) isEndPosition(position int) bool {
	return l.endPosition == position
}

func (l Ladder) GetType() UnitType {
	return LADDER
}
