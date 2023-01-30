// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
)

// Горутина заверщится при закрытии канала ch, который так же использентся для
// передачм данных
func gen() chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	c := gen()
	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
	}
	close(c)
}
