// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearch2(v []int, val int) (int, error) {

	left := 0
	right := len(v) - 1

	for left < right {

		mid := (left + right) / 2
		if val == v[mid] {
			return mid, nil
		}

		if val < v[mid] {
			right = mid - 1 // сдвиг правой границы
		}

		if val > v[mid] {
			left = mid + 1 // сдвиг левой границы
		}
	}

	if left == right {
		return left, nil
	}

	return -1, fmt.Errorf("error")
}

func main() {
	sorted := []int{-5, -3, -1, 0, 2, 3, 10, 15, 21}
	findInd := 8
	findNum := sorted[findInd]
	fmt.Printf("sorted[%d] = %d\n", findInd, findNum)

	ind, err := binarySearch2(sorted, 12)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Printf("index of %d is %d", findNum, ind)
}
