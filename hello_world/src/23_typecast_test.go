package try_test

import "testing"

type MyInt int64

func TestTypeCast(t *testing.T) {
	var a int32 = 100
	var b int64
	//不允许隐式类型转换
	// b = a //Cannot use 'b'(type int64) as type int64 in assignment
	b = int64(a)
	var c MyInt
	// c = b //Cannot use 'b'(type int64) as type MyInt in assignment
	c = MyInt(b)
	t.Log("a == ", a, ", b == ", b, ",c == ", c)
}