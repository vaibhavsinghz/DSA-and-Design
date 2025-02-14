package main

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return c.radius * 3.14 * c.radius
}

func NewCircle(radius int) IShape {
	return &Circle{radius: float64(radius)}
}
