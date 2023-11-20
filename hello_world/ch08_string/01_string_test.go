package try_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认零值""
	s = "hello"
	t.Log(len(s))
	// s[2] = '3' // string是不可变的 byte slice

	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	t.Log(s)
	s = "中"
	t.Log(s, len(s)) // 是byte个数

	// 将字符串s,转换为rune切片
	c := []rune(s)
	t.Log("len(c) == ", len(c))
	// t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode: %X", c[0])
	t.Logf("中 UTF8: %X", s)
}
