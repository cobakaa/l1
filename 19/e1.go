// Разработать программу, которая переворачивает подаваемую
//на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func reverseString(str string) string {
	v := []rune(str)
	n := len(v)

	for i := 0; i < n/2; i++ {
		v[i], v[n-1-i] = v[n-1-i], v[i]
	}

	return string(v)
}

func main() {
	var s string
	fmt.Scanln(&s)

	fmt.Println(reverseString(s))
}
