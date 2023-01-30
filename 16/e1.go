// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func quicksort(v []int) {

	if len(v) <= 1 {
		return
	}

	n := len(v) / 2

	seed := v[n]
	left := 0
	right := len(v) - 1
	for left <= right {
		for v[left] < seed { // нахождение элемента слева, который больше выбранного
			left++
		}

		for v[right] > seed { // нахождение элемента справа, который меньше выбранного
			right--
		}

		if left <= right {
			v[left], v[right] = v[right], v[left]
			left++
			right--
		}
	}

	quicksort(v[left:]) // сортировка правой половины

	quicksort(v[:right+1]) // сортировка левой половины
}

func main() {
	//v := []int{-10, 0, -3, 5, 5, 3, 2, 1, 6, 3, 1, 2, 4, 3245, 3, 3453, 2, -123}
	v := make([]int, 0)
	length := 100000
	for i := 0; i < length; i++ {
		v = append(v, rand.Int()) //Initialization array, length Length = 10000
	}

	now := time.Now()
	quicksort(v)
	fmt.Println("Array length:", length,
		"  time consuming:", time.Since(now).Microseconds(), sort.IntsAreSorted(v))
	//fmt.Println(v)
}
