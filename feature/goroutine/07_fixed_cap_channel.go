package main

import (
	"fmt"
	"time"
)

func TestChannel() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("push ", i)
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

//注意；模拟一个慢速的消费者，当生产者发现 缓冲channel满是，则阻塞等待
func RangeTravelChannel() {
	ch := make(chan int, 2)
	go Producer(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	TestChannel()
	RangeTravelChannel()
}
