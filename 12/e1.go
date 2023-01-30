// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func toSet(v []string) []string {
	tmp := make(map[string]struct{})
	res := make([]string, 0)
	for _, s := range v {
		if _, ok := tmp[s]; !ok {
			tmp[s] = struct{}{}
			res = append(res, s)
		}
	}

	return res
}

func main() {
	v := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(toSet(v))
}
