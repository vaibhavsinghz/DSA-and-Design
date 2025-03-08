package models

type Jump struct {
	Start, End int
}

func NewJump(start int, end int) *Jump {
	return &Jump{start, end}
}
