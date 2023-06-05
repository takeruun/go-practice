package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	// defer wg.Done() でもあり
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	// wg.Add(1)は、wg.Done()が1回呼ばれるまで待つという意味
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait()

	// 各メソッドのtime.Sleepをコメントアウトして実行すると、
	// goroutineの方も出力されるようになる
	// なせなら、goroutineの方が実行される前にmainが終了してしまうから
	// それを防ぐために、time.Sleepを入れている
	// time.Sleep(2000 * time.Millisecond)
}
