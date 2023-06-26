package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 1, 4}
	s := []string{"f", "e", "d", "c", "b"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Susi", 1},
		{"Aril", 5},
		{"Budi", 3},
		{"Mike", 2},
	}

	fmt.Println(i, s, p)

	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool {
		return p[i].Name < p[j].Name
	})
	fmt.Println(i, s, p)
}
