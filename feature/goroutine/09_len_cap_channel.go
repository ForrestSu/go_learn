package main

import "fmt"

/** go 语言还能发现运行时-死锁
 fatal error: all goroutines are asleep - deadlock!
 */
func main() {
	ch := make(chan string, 3)
	ch <- "alice"
	ch <- "bob"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value:", <-ch)
	fmt.Println("new length is", len(ch))
}
