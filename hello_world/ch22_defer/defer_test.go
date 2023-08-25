package main

import (
	"fmt"
	"testing"
)

// 4, 3, 2, 1
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	fmt.Println("4")
}
