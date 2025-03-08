package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	rng  *rand.Rand
	once sync.Once
)

func initRNG() {
	once.Do(func() {
		source := rand.NewSource(time.Now().UnixNano())
		rng = rand.New(source)
	})
}

func RandInt(min, max int) int {
	if min > max {
		panic("invalid range: min is greater than max")
	}
	initRNG()
	return rng.Intn(max-min+1) + min
}
