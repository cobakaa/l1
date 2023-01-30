// Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	v := []rune(str)
	last := len(v)
	for i := len(v) - 1; i >= 0; i-- {
		if v[i] == ' ' {
			b.WriteString(string(v[i+1 : last]))
			b.WriteRune(' ')
			last = i
		}
	}
	b.WriteString(string(v[:last]))
	return b.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	fmt.Println(reverseWords(s))
}
