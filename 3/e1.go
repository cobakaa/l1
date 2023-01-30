// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов с использованием конкурентных вычислений.

package main

import "fmt"

func main() {
	var v = []int{2, 4, 6, 8, 10}

	c := make(chan int)

	sqr := func(x int) {
		c <- x * x
	}

	// Подсчет квадрата в отдельной горутине и отправка результата в канал
	for _, n := range v {
		go sqr(n)
	}

	// Чтение из канала len(v) записей и вычисление суммы
	sum := 0
	for i := 0; i < len(v); i++ {
		sum += <-c
	}

	fmt.Println(sum)
}