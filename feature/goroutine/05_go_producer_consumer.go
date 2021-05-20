package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("channel closed.")
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
