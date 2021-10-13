package main

import (
	"context"
	"log"
	"time"
)

// ctx 作用：用来做超时控制

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case val := <-ctx.Done():
				log.Println("监控退出，停止了...", val)
				return
			default:
				log.Println("goroutine监控中...")
				time.Sleep(3 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	log.Println("可以了，通知监控停止")
	cancel()
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
