package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)

	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%d\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%d\n", len(a), cap(a), a)

	b := make([]int, 0)
	// メモリの確保をせずにスライスを作成する
	var c []int
	fmt.Printf("len=%d cap=%d value=%d\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%d\n", len(c), cap(c), c)
}
