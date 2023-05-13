package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func calc(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r, r2 := add(10, 20)
	fmt.Println(r, r2)

	r3 := calc(100, 2)
	fmt.Println(r3)

	f := func() {
		fmt.Println("inner func")
	}
	f()

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
