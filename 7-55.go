package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(3 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go goroutine1(c1)
	go goroutine2(c2)

	// packet from 1 しか受信されない
	// for result := range c1 {
	// 	fmt.Println(result)
	// }
	// for result := range c2 {
	// 	fmt.Println(result)
	// }

	// packet from 1 と packet from 2 が交互に受信される
	for {
		select {
		case result := <-c1:
			fmt.Println(result)
		case result := <-c2:
			fmt.Println(result)
		}
	}
}
