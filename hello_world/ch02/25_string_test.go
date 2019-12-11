package ch02_test

import "testing"

//
func TestString(t *testing.T) {
	var s string
	if s == "" {
		t.Log("a == <" + s + ">")
	}
	// len == 0
	t.Log("len(s) == ", len(s))
}
