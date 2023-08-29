package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

func TestAll(t *testing.T) {
	TestGolangRate(t)
	TestUberRate(t)
}

func TestGolangRate(t *testing.T) {
	ctx := context.Background()
	rl := rate.NewLimiter(1000000, 1) // 每秒钟产生 r 个

	last := time.Now()
	for i := 0; i < 10; i++ {
		rl.Wait(ctx)
		cur := time.Now()
		fmt.Println("last", cur.Sub(last))
		last = cur
	}
}

func TestUberRate(t *testing.T) {
	rl := ratelimit.New(1000000, ratelimit.WithoutSlack) // 每秒钟产生 r 个

	last := time.Now()
	for i := 0; i < 10; i++ {
		rl.Take()
		cur := time.Now()
		fmt.Println("last", cur.Sub(last))
		last = cur
	}
}
