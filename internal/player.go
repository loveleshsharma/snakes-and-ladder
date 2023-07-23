package internal

type Player struct {
	id         int
	name       string
	posOnBoard int
}

func NewPlayer(id int, name string) Player {
	return Player{
		id:         id,
		name:       name,
		posOnBoard: 0,
	}
}

func (p *Player) getPosition() int {
	return p.posOnBoard
}

func (p *Player) setPosition(position int) {
	p.posOnBoard = position
}
