// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

package main

import (
	"fmt"
	"unicode/utf8"
)

var justString string

func createHugeString(n int32) string {
	q := make([]rune, 0)
	for i := 0; i < 120; i++ {
		q = append(q, n)
	}
	return string(q)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // индекс среза учитывает байты, а не руны
	fmt.Println(len(justString), "bytes")
	fmt.Println(utf8.RuneCountInString(justString), "runes")
	justString = string([]rune(v)[:100]) // срез строится от слайса рун
	fmt.Println(len(justString), "bytes")
	fmt.Println(utf8.RuneCountInString(justString), "runes")
}

func main() {
	someFunc()
}
