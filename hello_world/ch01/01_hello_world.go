package main

import (
	"fmt"
)

func main() {
	deadLock()
	// os.Exit(-1)
}

// deadLock is a deadlock function
func deadLock() {
	fmt.Println("hello world!")
	// deadlock!
	done := make(chan bool)
	<-done
}
