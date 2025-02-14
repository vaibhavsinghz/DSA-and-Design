package main

import (
	"fmt"
	"time"
)

func channelBlocks() {
	ch := make(chan int)
	ch <- 42 //block
	fmt.Println(<-ch)
}

func channelWork() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
}
func channelWork2() {
	ch := make(chan int, 2) //bufferred channel
	ch <- 49
	fmt.Println(<-ch)

	// ch <- 499
	// fmt.Println(<-ch)
}

func channelWork3() {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending : ", i)
			ch <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Received : ", <-ch)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// channelBlocks()
	// channelWork()
	// channelWork2()
	channelWork3()
}
