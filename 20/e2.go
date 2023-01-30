package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords2(str string) string {
	v := strings.Split(str, " ")
	n := len(v)
	for i := 0; i < n/2; i++ {
		v[i], v[n-1-i] = v[n-1-i], v[i]
	}

	return strings.Join(v, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	fmt.Println(reverseWords2(s))
}
