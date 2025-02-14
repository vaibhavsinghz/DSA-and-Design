package main

import "fmt"

func Use(shape IShape) {
	fmt.Println(shape.Area())
}
