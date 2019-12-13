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

	//将字符串s,转换为rune切片
	c := []rune(s)
	t.Log("len(c) == ", len(c))
	//t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode: %X", c[0])
	t.Logf("中 UTF8: %X", s)
}

// 字符串变量，会自动遍历里面的 rune
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
		//t.Logf("%c %x", c, c)
	}
	//这里 [1] 表示都要使用第一个参数
}
