package main

import (
	"fmt"
	"testing"
)

func hello() {
	fmt.Println("co: hello world goroutine")
}

// 启动一个协程，执行 hello()
func TestCo(t *testing.T) {
	go hello()
	fmt.Println("main: function")
}
