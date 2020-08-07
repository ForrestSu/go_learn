package ch03_compare_loop_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// Go语言 &^ 表示按bit置0
// 等价于c语言里的 &=~ (eg: a &= ~1;)
func TestBitClear(t *testing.T) {
	a := 7 //0111
	// remove read permission
	a = a &^ Readable
	// remove execute permission
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
