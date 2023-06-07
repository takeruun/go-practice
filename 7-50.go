package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// channel に送信
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// channel に送信
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15 になる
	go goroutine1(s, c)
	go goroutine2(s, c)

	// channel から受信
	x := <-c
	fmt.Println(x)

	y := <-c
	fmt.Println(y)
}
