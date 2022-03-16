package ch02_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Go语言里面，字符串为值类型
func TestString(t *testing.T) {
	var s string
	if s == "" {
		t.Log("a == <" + s + ">")
	}
	assert.True(t, len(s) == 0)
}
