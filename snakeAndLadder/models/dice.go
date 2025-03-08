package models

import "Self/snakeAndLadder/utils"

type Dice struct {
	min, max int
}

func NewDice1(min, max int) *Dice {
	return &Dice{min, max}
}

func (d *Dice) Roll() int {
	return utils.RandInt(d.min, d.max)
}
