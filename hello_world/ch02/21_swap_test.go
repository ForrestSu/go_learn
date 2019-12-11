package ch02_test

import "testing"

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	// swap a,b
	a, b = b, a
	t.Log("a == ", a, ", b == ", b)
}