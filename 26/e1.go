// Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

package main

import (
	"fmt"
	"unicode"
)

func isSymbolsUnique(str string) bool {
	m := make(map[rune]struct{})
	for _, c := range []rune(str) {
		c = unicode.ToLower(c)
		if _, ok := m[c]; ok {
			return false
		}
		m[c] = struct{}{}
	}

	return true
}

func main() {

	v := []string{"abcd", "abCdefAaf", "aabcd", "абВндА", "аб Внде"}

	for _, s := range v {
		fmt.Println(s+" -", isSymbolsUnique(s))
	}

}
