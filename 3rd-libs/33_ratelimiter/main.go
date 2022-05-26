// Package main Rate limiter demo
package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	l := rate.NewLimiter(10, 5)

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case tm := <-ticker.C:
			log.Println(tm)
		default:
			log.Println(l.Wait(ctx))
			time.Sleep(5 * time.Millisecond)
		}
	}
}
