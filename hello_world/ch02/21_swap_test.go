package ch02_test

import "testing"

// 交换两个变量值
func TestSwap(t *testing.T) {
	a := 1
	b := 2
	// swap a,b
	a, b = b, a
	t.Log("a == ", a, ", b == ", b)
}