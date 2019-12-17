package test_go_oo

import "testing"

// 接口定义 duck type
type Programmer interface {
	WriteHelloWorld() string
}

// struct
type GoProgrammer struct {
}

// 接口实现
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	// g 接口变量
	var g = GoProgrammer{}
	t.Log(g.WriteHelloWorld())
}
