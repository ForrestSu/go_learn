package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEmptySelect(t *testing.T) {
	// fatal error: all goroutines are asleep - deadlock!
	// select {}
}

func TestChannelSelect(t *testing.T) {
	ch01 := make(chan int)
	ch02 := make(chan int)
	go func(ch01 chan int, ch02 chan int) {
		fmt.Println("enter co-routine...")
		select {
		case rc1 := <-ch01:
			fmt.Println("rc1", rc1)
		case rc2 := <-ch02:
			fmt.Println("rc1", rc2)
		default:
			fmt.Println("no data")
		}
	}(ch01, ch02)

	time.Sleep(time.Second)
	close(ch01)
	// panic: send on closed channel
	ch01 <- 100
}

// 采用 select 实现超时
func TestSelectTimeOut(t *testing.T) {
	msgCh := make(chan int)
	go func() {
		fmt.Println("enter co-routine...")
		isRun := true
		for isRun {
			select {
			case rc1, ok := <-msgCh:
				fmt.Println("rc1", rc1)
				if ok == false {
					isRun = false
					fmt.Println("channel closed")
				}
			case rc2 := <-time.After(time.Millisecond * 200):
				fmt.Println("time out 1 sec", rc2)
			}
		}
		fmt.Println("exit co-routine...")
	}()
	time.Sleep(time.Second * 2)
	// 2秒之后发送一个消息
	msgCh <- 10
	close(msgCh)
}
