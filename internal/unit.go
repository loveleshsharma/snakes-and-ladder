package internal

type UnitType int

const (
	SNAKE UnitType = iota
	LADDER
)

type Unit interface {
	GetType() UnitType
}
