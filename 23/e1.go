// Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
	"math/rand"
)

func removeElement1(v []int, i int) []int {
	return append(v[:i], v[i+1:]...) // Вставка в срез до iго элемента того, что стоит после него
}

func removeElement2(v []int, i int) []int {
	copy(v[i:], v[i+1:]) // сдвиг элементов после iго на 1 влево и удаление последнего
	return v[:len(v)-1]
}

func main() {
	v := make([]int, 0)
	for i := 0; i < 100; i++ {
		v = append(v, rand.Intn(10))
	}
	fmt.Println(v, len(v))
	v = removeElement1(v, 2)
	fmt.Println(v, len(v))
	v = removeElement2(v, 0)
	fmt.Println(v, len(v))
}
