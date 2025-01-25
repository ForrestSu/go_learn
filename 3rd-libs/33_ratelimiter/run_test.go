package main

import (
	"context"
	"log"
	"testing"

	"golang.org/x/time/rate"
)

// TestPer100 100 QPS
func TestPer100(t *testing.T) {
	run()
}

func run() {
	ctx := context.Background()
	// 计算每秒允许的请求速率
	rateLimit := rate.Limit(1)
	// 设置令牌桶的最大容量，这里设置为 200，意味着可以在瞬间处理 200 个请求（如果之前有积攒的令牌）
	burst := 2
	// 创建一个新的速率限制器
	limiter := rate.NewLimiter(rateLimit, burst)
	// 模拟持续发送查询请求
	k := 0
	for {
		if err := limiter.Wait(ctx); err != nil {
			log.Println("等待令牌时出错:", err)
			return
		}
		Query(k)
		k++
	}
}

// Query 模拟查询请求的函数
func Query(k int) {
	log.Printf("执行查询请求 %d\n", k)
}
