package ch02_test

import "testing"

type MyInt int64

// (1) Go 语言不支持隐式类型转换，必须显示转换
// (2) 别名和原有类型也不能进行隐式类型转换

func TestTypeCast(t *testing.T) {
	var a int32 = 100
	var b int64
	// 不允许隐式类型转换
	// b = a //Cannot use 'b'(type int64) as type int64 in assignment
	b = int64(a)
	var c MyInt
	// c = b //Cannot use 'b'(type int64) as type MyInt in assignment
	c = MyInt(b)
	t.Log("a == ", a, ", b == ", b, ",c == ", c)
}
