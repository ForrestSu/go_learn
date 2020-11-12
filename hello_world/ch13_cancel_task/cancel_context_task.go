package main

import (
	"context"
	"fmt"
	"time"
)

// 第三层的子任务
func subTask(ctx context.Context, name string, node string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name+"-"+node, " exit...")
			return
		case <-time.After(time.Second):
			fmt.Println(name+"-"+node, "time out")
		}
	}
}

//ctx以参数传递
func work(ctx context.Context, name string, node string) {
	// 继续创建子任务
	subCtx, _ := context.WithCancel(ctx)
	go subTask(subCtx, "level3-"+node, "left")
	go subTask(subCtx, "level3-"+node, "right")
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name+"-"+node, " exit...")
			return
		case <-time.After(time.Second):
			fmt.Println(name+"-"+node, "time out")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "level2", "A")
	go work(ctx, "level2", "B")

	fmt.Println("do cancel")
	// 通知ctx创建的协程，全部退出
	time.Sleep(time.Second * 2)

	fmt.Println("do cancel...")
	cancel()
	fmt.Println("cancel sent...")
}
