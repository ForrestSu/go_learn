package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func commaOK(v interface{}) *int {
	val, _ := v.(*int)
	return val
}

func NotCommaOK(v interface{}) *int {
	val := v.(*int)
	return val
}

func TestPass(t *testing.T) {
	// 正常返回 nil
	assert.Nil(t, commaOK(nil))
	// 必然 panic
	assert.Panics(t, func() {
		NotCommaOK(nil)
	})
}
