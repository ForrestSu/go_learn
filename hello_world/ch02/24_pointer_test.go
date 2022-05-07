package ch02_test

import "testing"

// Go 语言支持指针类型，但是指针不支持运算
func TestPointer(t *testing.T) {
	var a int = 10
	aPtr := &a
	// aPtr = aPtr + 1 // error
	t.Log("a == ", a, ", aPtr == ", aPtr)
	t.Logf("%T %T", a, aPtr)
}
