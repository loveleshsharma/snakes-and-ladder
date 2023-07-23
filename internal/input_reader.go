package internal

type InputReader interface {
	Read() ([]Snake, []Ladder, []Player, error)
}
