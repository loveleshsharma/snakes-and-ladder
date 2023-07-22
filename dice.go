package main

import (
	"math/rand"
)

type Dice struct {
}

func NewDice() Dice {
	return Dice{}
}

func (d Dice) roll() int {
	return rand.Intn(5) + 1
}
