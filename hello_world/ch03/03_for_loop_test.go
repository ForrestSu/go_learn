package ch03_test

import (
	"testing"
	"time"
)

// test for loop
func TestFor(t *testing.T) {
	n := 0
	/* while( n < 5) */
	for n < 5 {
		n++
		t.Log("n == ", n)
	}
}

// 测试死循环
func TestForNoDie(t *testing.T) {
	for {
		println("hello")
		time.Sleep(1000)
	}
}
