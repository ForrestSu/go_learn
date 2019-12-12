package ch03_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

// Go语言 按照bit位, 置0
func TestBitClear(t *testing.T) {
	a := 7 //0111
	// remove read permission
	a = a &^ Readable
	// remove execute permission
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
