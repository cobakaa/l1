package main

import "fmt"

func main() {
	// В случае закрытия канала quit горутина остановится
	quit := make(chan struct{})
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println(<-ch)
			}
		}
	}()
	for i := 0; i < 16; i++ {
		ch <- i
	}

	close(quit)
}
