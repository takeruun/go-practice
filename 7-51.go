package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	// fatal error: all goroutines are asleep - deadlock!
	// ch はバッファが 2 なので、2 つまでしか値を送信できない
	// ch <- 300
	// fmt.Println(len(ch))

	// range すると、ch が close されるまで値を受信し続ける
	// close しないと、fatal error: all goroutines are asleep - deadlock!
	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
