package main

import (
	"fmt"
)

func main() {
	var test string
	if test == "" {
		fmt.Println("string test is empty !")
	}
	fmt.Println("hello world!")
	// deadlock!
	done := make(chan bool)
	<-done
}
