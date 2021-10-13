package main

import (
	"fmt"
	"time"
)

func helloFunc(done chan int) {
	fmt.Println("enter...")
	time.Sleep(3 * time.Second)
	fmt.Println("Hello world goroutine")
	done <- 1
}

func main() {
	done := make(chan int)
	go helloFunc(done)
	notify := <-done
	fmt.Println("main function ", notify, len(done))
}
