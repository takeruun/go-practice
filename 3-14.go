package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	/*
		var b [2]int = [2]int{100, 200}
		// 配列個数指定したものには append できない
		b = append(b, 300)
	*/

	// 配列個数指定しないものには append できる
	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)
}
