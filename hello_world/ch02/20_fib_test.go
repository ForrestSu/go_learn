package ch02_test

import "testing"

// 计算斐波那契数列
func TestFibonacci(t *testing.T) {
	var a int = 1
	var b int = 2
	for i := 0; i <= 5; i++ {
		c := a + b
		a = b
		b = c
		t.Log("c = ", c)
	}
}
