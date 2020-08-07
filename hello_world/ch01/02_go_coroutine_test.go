package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	fmt.Println("hello world!")
	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU()) //逻辑核心数 12
	fmt.Println(runtime.NumGoroutine())//1
	fmt.Println(runtime.NumCgoCall())//33
	go func() {
		fmt.Println("go routine count:",runtime.NumGoroutine()) //3
	}()
	go func() {
		fmt.Println("go routine count2:",runtime.NumGoroutine()) //4
	}()
	time.Sleep(time.Millisecond)
}