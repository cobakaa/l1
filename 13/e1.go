// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
