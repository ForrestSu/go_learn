package try_test

import (
	"strconv"
	"strings"
	"testing"
)

// 字符串常用函数
func TestStringFN(t *testing.T) {
	s := "A,B,C"
	// 1 拆分字符串
	parts := strings.Split(s, ",")

	t.Log(len(parts))
	for _, part := range parts {
		t.Log(part)
	}
	// 字符串连接
	t.Log(strings.Join(parts, "-"))
}

// 字符串和整数转换
func TestStringConvert(t *testing.T) {
	// 1 整数 => 字符串
	var s string
	s = strconv.Itoa(67)
	t.Logf("type: %[1]T, %[1]s", s)

	// 2 字符串转为整数
	intVar, _ := strconv.Atoi(s)
	t.Log(intVar)
}
