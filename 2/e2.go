package main

import "fmt"

func main() {
	var v = []int{2, 4, 6, 8, 10}

	c := make(chan int)

	// Отправка в канал результата вычислений
	sqr := func(x int) {
		c <- x * x
	}

	// Подсчет квадрата в отдельной горутине и отправка результата в канал
	for _, n := range v {
		go sqr(n)
	}

	// Чтение из канала len(v) записей для избежания deadlock
	for i := 0; i < len(v); i++ {
		fmt.Println(<-c)
	}
}
