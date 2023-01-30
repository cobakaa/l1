// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

// Отправка из слайса в первый канал

func toFirstCh(v []int) <-chan int {
	first := make(chan int)
	go func() {
		for _, n := range v {
			first <- n
		}
		close(first)
	}()
	return first
}

// Отправка из переданного канала в следующий

func toSecondCh(ch <-chan int) <-chan int {
	second := make(chan int)
	go func() {
		for n := range ch {
			second <- n * 2
		}
		close(second)
	}()
	return second
}

func main() {

	v := make([]int, 0)

	for i := 0; i < 10; i++ {
		v = append(v, i)
	}

	// Вывод полученных чисел
	for s := range toSecondCh(toFirstCh(v)) {
		fmt.Println(s)
	}

}
