// Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.

package main

import (
	"fmt"
	"l1/24/point"
)

func main() {
	p1 := point.NewPoint(0, 0)
	p2 := point.NewPoint(3, 4)
	fmt.Println(p1.Distance(p2))
	fmt.Println(p2.Distance(p1))
}
