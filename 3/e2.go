package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	v := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var sum int64 = 0

	sqr := func(x int) {
		// сигнал, что элемент группы завершил свое выполнение
		defer wg.Done()
		atomic.AddInt64(&sum, int64(x*x))
	}

	wg.Add(len(v))
	for _, n := range v {
		go sqr(n)
	}
	wg.Wait()
	fmt.Println(sum)
}