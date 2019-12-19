package ch11_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {

	//Go断言: 1 直接判断是否为 int 类型
	//value, ok := p.(int)

	// 2 配合switch使用，可使用 p.(type)
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknown", v)
	}
}

func TestFn(t *testing.T) {
	DoSomething(10)
	DoSomething("Alice")
	DoSomething(12.9)
}
