// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Border struct {
	left  int
	right int
}

type Stack struct {
	stack []Border
}

func (st *Stack) Push(b Border) {
	st.stack = append(st.stack, b)
}

func (st *Stack) Top() (Border, error) {
	if len(st.stack) == 0 {
		return Border{}, errors.New("empty")
	}
	return st.stack[len(st.stack)-1], nil
}

func (st *Stack) Pop() {
	size := len(st.stack)
	if size > 0 {
		st.stack = st.stack[:size-1]
	}
}

func (st *Stack) Empty() bool {
	return len(st.stack) == 0
}

func qsort(v []int) { // нерекурсивный вариант quicksort

	if len(v) <= 1 {
		return
	}

	st := Stack{make([]Border, 0)}
	b := Border{0, len(v) - 1}
	st.Push(b)
	for !st.Empty() {
		t, _ := st.Top()
		left := t.left
		right := t.right
		st.Pop() // выбор запроса из стека
		for left < right {
			i := left
			j := right
			x := v[(left+right)/2]
			for i <= j {
				for v[i] < x {
					i++
				}
				for v[j] > x {
					j--
				}
				if i <= j {
					v[i], v[j] = v[j], v[i]
					i++
					j--
				}
			}
			if i < right { // откладывание запроса на сортировку правой части в стек
				t.left = i
				t.right = right
				st.Push(t)
			}
			right = j // обновление ограничения леврй части
			// после этого она сортируется сразу
		}
	}
}

func main() {
	//v := []int{-10, 0, -3, 5, 5, 3, 2, 1, 6, 3, 1, 2, 4, 3245, 3, 3453, 2, -123}
	v := make([]int, 0)

	for i := 0; i < 100000; i++ {
		v = append(v, rand.Int()) //Initialization array, length Length = 10000
	}

	length := len(v)
	now := time.Now()
	qsort(v)
	fmt.Println("Array length:", length,
		"  time consuming:", time.Since(now).Microseconds(), sort.IntsAreSorted(v))
	//fmt.Println(v)
}
