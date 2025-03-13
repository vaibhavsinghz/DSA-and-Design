package main

import (
	"math/rand"
	"time"
)

const (
	Stone = iota
	Paper
	Scissor
)

func main() {}

func RandInt(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return rng.Intn(max-min+1) + min
}
