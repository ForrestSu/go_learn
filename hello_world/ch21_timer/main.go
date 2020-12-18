package main

import (
	"context"
	"log"
	"time"
)

func TestTimer(ctx context.Context) {
	multiAidsTicker := time.NewTicker(3 * time.Second)
	go func() {
		log.Println("timer...")
		for {
			select {
			case <-ctx.Done():
				log.Println("done ...")
				return
			case <-multiAidsTicker.C:
				log.Println("timer...")
			}
		}
	}()
	timer2 := time.NewTicker(3 * time.Second)
	go func() {
		log.Println("timer2...")
		for {
			select {
			case <-ctx.Done():
				timer2.Stop()
				log.Println("done2 ...")
				return
			case <-timer2.C:
				log.Println("timer2...")
			}
		}
	}()
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		log.Println("cancel ...")
	}()
	TestTimer(ctx)

	log.Println("will sleep...")
	time.Sleep(15 * time.Second)
}
