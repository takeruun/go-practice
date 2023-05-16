package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo(10)
	foo(10, 11)
	foo(10, 11, 12)

	s := []int{1, 2, 3}
	foo(s...)
}
