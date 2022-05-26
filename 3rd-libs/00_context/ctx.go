package main

import (
	"context"
	"log"
	"time"
)

// context 主要作用是用来做 Cancel
// context 本质是一个 tree 的结构，当前节点的cancel信号会传递到其下的所有子节点

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
