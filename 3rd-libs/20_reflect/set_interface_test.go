package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Student struct {
	Name string
	Age  int
}

// 反射设置值
func setInterfaceValue(fromPtr interface{}, toPtr interface{}) {
	reflect.ValueOf(toPtr).Elem().Set(reflect.ValueOf(fromPtr).Elem())
}

func TestReflectSet(t *testing.T) {
	from := &Student{
		Name: "Bob",
		Age:  18,
	}
	to := &Student{}
	setInterfaceValue(from, to)
	assert.Equal(t, from, to)
}
