package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func runTimer(ctx context.Context, id int) {
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		log.Printf("enter => [%d]...\n", id)
		for {
			select {
			case <-ctx.Done():
				log.Printf("done[%d] ...\n", id)
				return
			case <-ticker.C:
				log.Printf("timer[%d] ...\n", id)
			}
		}
	}()
}

func TestMultiTimer(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		log.Println("cancel ...")
	}()
	runTimer(ctx, 0)
	runTimer(ctx, 1)
	log.Println("will sleep...")
	time.Sleep(10 * time.Second)
}
