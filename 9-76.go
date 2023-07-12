package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// 並行処理数の制限 1 は 1 つの処理しか実行できない
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// これは、1 つの処理しか実行できない
	// s.Acquire　は処理を止めるけど、s.TryAcquire は処理を止めずに、処理を続ける
	// return で処理を終了させる

	// isAcquire := s.TryAcquire(1)
	// if !isAcquire {
	// 	fmt.Println("Could not get lock")
	// 	return
	// }
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)

	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
