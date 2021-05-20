package main

import (
	"fmt"
	"time"
)

func hello(done chan int) {
	fmt.Println("enter...")
	time.Sleep(3 * time.Second)
	fmt.Println("Hello world goroutine")
	done <- 1
}

func main() {
	done := make(chan int)
	go hello(done)
	notify := <-done
	fmt.Println("main function ", notify, len(done))
}
