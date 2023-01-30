package main

import "fmt"

func main() {
	// Проверка на то, что канал не закрыт
	ch := make(chan int)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}()

	for i := 0; i < 8; i++ {
		ch <- i
	}

	close(ch)
}
