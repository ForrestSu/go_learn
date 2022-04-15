package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Filter 过滤器接口
type Filter interface {
	Watch() <-chan EventType
}

// 实现消息通知
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1 启动任务
	f := newNotifier()
	mockPolling(ctx, f)
	go func() {
		for event := range f.Watch() {
			log.Printf("reciving event: %v", event)
		}
		log.Printf("event channel closed!")
	}()

	// waiting signal
	fmt.Println("---")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("exit")
}

func mockPolling(ctx context.Context, d *Notifier) {
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("timer...")
				d.Send(EventDataChanged)
				// if data changed
			case <-ctx.Done():
				log.Println("polling done")
				return
			}
		}
	}()
}
