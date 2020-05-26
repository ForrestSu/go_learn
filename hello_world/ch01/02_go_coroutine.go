package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("hello world!")
	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU()) //4
	fmt.Println(runtime.NumGoroutine())//1
	fmt.Println(runtime.NumCgoCall())//0
	go func() {
		fmt.Println("go routine count:",runtime.NumGoroutine()) //3
	}()
	go func() {
		fmt.Println("go routine count2:",runtime.NumGoroutine()) //3
	}()
	time.Sleep(time.Millisecond)
}