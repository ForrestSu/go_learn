package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

func Consumer0(ch chan int) {
	iCount := 0
	for {
		data, ok := <-ch
		fmt.Println("ch0:", data)
		if ok == false {
			fmt.Println("channel 0 closed! iCount=", iCount)
			break
		}
		iCount++
	}
	wg.Done()
}

func Consumer1(ch chan int) {
	iCount := 0
	for {
		data, ok := <-ch
		fmt.Println("ch1:", data)
		if ok == false {
			fmt.Println("channel 1 closed! iCount=", iCount)
			break
		}
		iCount++
	}
	wg.Done()
}

func TestMultiConsumer(t *testing.T) {
	ch := make(chan int)
	wg.Add(2)
	go Consumer0(ch)
	go Consumer1(ch)

	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 1)
	}
	fmt.Println("main wait")
	// 广播,告诉其他接收者退出
	close(ch)
	wg.Wait()
}
