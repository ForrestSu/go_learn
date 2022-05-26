package main

import "fmt"

/** go 语言还能发现运行时-死锁
fatal error: all goroutines are asleep - deadlock!
*/
func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
