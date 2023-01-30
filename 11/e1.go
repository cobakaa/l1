// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersect(s1, s2 []int) []int {
	tmp := make(map[int]struct{}) // аналог множества
	res := make([]int, 0)
	for _, n := range s1 {
		if _, ok := tmp[n]; !ok {
			tmp[n] = struct{}{}
			res = append(res, n)
		}
	}

	for _, n := range s2 {
		if _, ok := tmp[n]; !ok {
			tmp[n] = struct{}{}
			res = append(res, n)
		}
	}

	return res
}

func main() {
	s1 := []int{1, 2, 5, 5, 6, 3, 2}
	s2 := []int{2, 3, 4, 4, 1, 9}

	fmt.Println(intersect(s1, s2))
}
