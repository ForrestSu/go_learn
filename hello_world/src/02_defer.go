package main

import "fmt"

func main() {

	fmt.Println("print 0")
    // push stack
	defer fmt.Println("print 1")
    // push stack
	defer fmt.Println("print 2")

	fmt.Println("print 3")
}