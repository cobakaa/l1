// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

func binarySearch(v []int, val int) (int, error) {
	mid := len(v) / 2

	if val == v[mid] {
		return mid, nil
	}

	if len(v) <= 1 {
		return -1, fmt.Errorf("error")
	}

	if val < v[mid] {
		return binarySearch(v[:mid], val) // поиск в левой части
	}

	ind, err := binarySearch(v[mid:], val) // поиск в правой части
	return ind + mid, err

}

func main() {
	sorted := []int{-5, -3, -1, 0, 2, 3, 10, 15, 21, 123}
	findInd := 6
	findNum := sorted[findInd]
	fmt.Printf("sorted[%d] = %d\n", findInd, findNum)

	ind, err := binarySearch(sorted, findNum)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Printf("index of %d is %d", findNum, ind)
}
