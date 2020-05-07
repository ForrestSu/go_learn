package test_go_oo

import (
	"fmt"
	"testing"
)

type Human struct {
	name  string
	age   int
	email string
}

type Student struct {
	//h      Human //has-a 非匿名组合
	Human //is-a(Pseudo) 匿名组合
	school string
	//name string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am [%s] you can contacts me on [%s]\n", h.name, h.email)
}

func (s *Student) SayHi() {
	fmt.Printf("Hi student, I am [%s], you can contacts me on [%s]\n", s.Human.name, s.Human.email)
}

func TestIsA(t *testing.T) {
	stu := Student{Human{"Alice", 25, "123@qq.com"}, "MIT"}
	t.Log(stu)
	stu.Human.SayHi()
	stu.SayHi()
	t.Log(stu.name)
	t.Log(stu.Human.name)
}