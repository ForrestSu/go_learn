// Package main Rate limiter demo
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	testWait() // blocking
	// testReserve()
	// testAllow() // non-blocking
}

func testWait() {
	// 每秒钟1个
	lim := rate.NewLimiter(rate.Every(time.Second), 1)
	for i := 0; i < 5; i++ {
		if err := lim.WaitN(context.Background(), 1); err != nil {
			log.Println(err)
		}
		fmt.Printf("%v wait done\n", nowMs())
	}
}

func testReserve() {
	// 每秒钟1个
	lim := rate.NewLimiter(rate.Every(time.Second), 1)
	for i := 0; i < 5; i++ {
		rev := lim.ReserveN(time.Now(), 1)
		if !rev.OK() {
			// Not allowed to act! Did you remember to set lim.burst to be > 0 ?
			panic("err")
		}
		delay := rev.Delay() // 等待时间
		fmt.Printf("%s i=%d, delay=%v\n", nowMs(), i, delay)
		time.Sleep(delay)
	}
}

// Allow returns true if the event may happen at time now.
func testAllow() {
	// 每秒钟1个
	lim := rate.NewLimiter(rate.Every(time.Second), 1)
	for i := 0; i < 5; i++ {
		if lim.Allow() {
			log.Println("allow")
		}
	}
}

func nowMs() string {
	return time.Now().Format("15:04:05.000")
}
