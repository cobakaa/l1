// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := make([]interface{}, 0)
	v = append(v, 42)                 // int
	v = append(v, "str")              // string
	v = append(v, true)               // bool
	v = append(v, make(chan rune))    // chan rune
	v = append(v, make(chan float64)) // chan float64

	for _, x := range v {
		switch reflect.ValueOf(x).Kind() { // для того, чтобы распознавать любой канал как канал
		case reflect.Int:
			{
				fmt.Println("int")
			}

		case reflect.String:
			{
				fmt.Println("string")
			}

		case reflect.Bool:
			{
				fmt.Println("bool")
			}

		case reflect.Chan:
			{
				fmt.Println("channel")
			}
		}
	}
}
