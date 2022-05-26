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

func runTask(ctx context.Context, secTO int, n *int) {
	log.Printf(">>> task: %d ...\n", *n)

	select {
	case <-ctx.Done():
		log.Println("ctx cancel!")
	case <-time.After(time.Duration(secTO) * time.Second):
		log.Printf("secTO = %d, is time out!\n", secTO)
	}

	log.Printf("<<< task: %d Ok.\n", *n)
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	taskList := []int{3, 8, 3}
	taskCnt := len(taskList)

	wg := sync.WaitGroup{}
	wg.Add(taskCnt)
	for _, task := range taskList {
		// 并发执行多个任务
		go func(ctx1 context.Context, n int) {
			defer wg.Done()
			counter := 0
			runTask(ctx1, n, &counter)
			if ctx1.Err() != nil {
				log.Printf("n == %d, ctx err=%+v\n", n, ctx1.Err())
				if ctx1.Err() == context.DeadlineExceeded {
					log.Println("equal DeadlineExceeded")
				}
			}
		}(ctx, task)
	}
	wg.Wait()
	log.Println("exit...")
}
