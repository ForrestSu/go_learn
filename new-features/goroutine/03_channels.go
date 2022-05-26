package main

import (
	"fmt"
)

// 创建一个 int 信道 channel
func TestCreatChannel() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		// <chan int> len = 0
		fmt.Printf("Type of a is <%T> %d", a, len(a))
	}
}
func main() {
	TestCreatChannel()
}
