package main

import (
	"context"
	"fmt"
	"time"
)

func temp(ctx context.Context) {
	i := 0
	for {
		i++
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("sending", i)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go temp(ctx)
	time.Sleep(1 * time.Second)
	cancel()
}
