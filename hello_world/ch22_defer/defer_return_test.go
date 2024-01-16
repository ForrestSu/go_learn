package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fn() (n int) {
	defer func() {
		n = 100
	}()
	return n
}

func ret() *student {
	s := &student{Name: "Alice"}
	defer func() {
		s = &student{Name: "Bob"}
	}()
	return s
}

// 命名出参方式：defer修改返回指针，会生效
func retVar() (s *student) {
	s = &student{Name: "Alice"}
	defer func() {
		s = &student{Name: "Bob"}
	}()
	return
}

func TestFib(t *testing.T) {
	assert.Equal(t, 100, fn())
	t.Logf("name = %s", ret().Name)    // output: name = Alice
	t.Logf("name = %s", retVar().Name) // output: name = Bob
}
