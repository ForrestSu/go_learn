package main

import (
	"fmt"
	"testing"
)

// recover 只能捕获当前协程的 panic
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("has panic() = %+v!\n", err)
		} else {
			fmt.Println("no panic()!")
		}
	}()
	panic("hhh")
}

// 无法捕获
func TestNotRecover(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("has panic() = %+v!\n", err)
		} else {
			fmt.Println("no panic()!")
		}
	}()

	go func() {
		panic("hhh")
	}()
}
