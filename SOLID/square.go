package main

type Square struct {
	size float64
}

func (c *Square) Area() float64 {
	return c.size * c.size
}

func NewSquare(radius int) IShape {
	return &Square{size: float64(radius)}
}
