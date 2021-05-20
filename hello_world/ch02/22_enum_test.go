package ch02_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 测试枚举常量，以及 bit mask 常量
func TestEnum(t *testing.T) {
	t.Log("Monday == ", Monday, ", Tuesday == ", Tuesday)
	t.Log("Readable == ", Readable, ", Executable == ", Executable)
}

// 测试位运算
func TestBit(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable, a&Writable, a&Executable)
}
