package ch11_duck_type

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() string {
	return "System.out.Println(\"Hello World\")"
}

//Duck type
func PrintHelloWorld(p Programmer) {
	fmt.Printf("%T, %s\n", p, p.WriteHelloWorld())
}
func TestDuckType(t *testing.T) {
	g := new(GoProgrammer)
	t.Log(g.WriteHelloWorld())

	j := new(JavaProgrammer)
	t.Log(j.WriteHelloWorld())

	//duck type, 入参数只能为Pointer
	PrintHelloWorld(g)
	PrintHelloWorld(j)
}
