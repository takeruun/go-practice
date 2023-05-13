package main

import "fmt"

func incrementGenerator() func() int {
	// x がクロージャーになる
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleAreaGenerator(pi float64) func(radius float64) float64 {
	// pi がクロージャーになる
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	// increment := func() int {
	// 	x++
	// 	return x
	// }

	// fmt.Println(increment())
	// fmt.Println(increment())
	// fmt.Println(increment())
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleAreaGenerator(3.14)
	fmt.Println(c1(2))

	c2 := circleAreaGenerator(3)
	fmt.Println(c2(2))
}
