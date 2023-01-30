package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// С помощию контекста и отправки сигнала cancel()
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				i++
				fmt.Println(i)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}(ctx)

	time.Sleep(600 * time.Millisecond)
	cancel()
}
