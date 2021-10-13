package ch19_test_assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithAssert(t *testing.T) {
	var actual = 11
	var expect = 12
	// 使用断言比较相等
	assert.NotEqualf(t, expect, actual, "occur error!")
	t.Log("exit!")
}
