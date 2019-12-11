package ch03_test

import "testing"

// 比较数组
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 5}
	b := [...]int{1, 2, 3, 5}

	// equals
	if a == b {
		t.Log("a == b")
	} else {
		t.Log("a != b")
	}

	// not equals
	c := [...]int{1, 2, 3, 5}
	d := [...]int{5, 2, 3, 1}

	if c == d {
		t.Log("c == d")
	} else {
		t.Log("c != d")
	}
}
