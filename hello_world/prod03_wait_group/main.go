package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func runTask(ctx context.Context, n int) {
	log.Printf("task: %d ...\n", n)

	time.Sleep(3 * time.Second)
	log.Printf("task: %d Ok.\n", n)
}

func main() {

	ctx := context.Background()

	taskList := []int{10, 20, 30, 50}
	taskCnt := len(taskList)

	wg := sync.WaitGroup{}
	wg.Add(taskCnt)
	for _, task := range taskList {
		// 并发执行多个任务
		go func(ctx1 context.Context, n int) {
			defer wg.Done()
			runTask(ctx1, n)
		}(ctx, task)
	}
	wg.Wait()
	log.Println("exit...")
}
