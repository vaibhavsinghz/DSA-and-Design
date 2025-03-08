package models

type Cell struct {
	Jump *Jump
}

func NewCell() *Cell {
	return &Cell{}
}

func (c *Cell) SetJump(j *Jump) {
	c.Jump = j
}

func (c *Cell) GetJump() *Jump {
	return c.Jump
}
