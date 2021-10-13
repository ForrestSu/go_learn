package ch03_compare_loop_test

import (
	"testing"
)

// test for loop
func TestFor(t *testing.T) {
	n := 0
	/* while( n < 5) */
	for n < 5 {
		n++
		t.Log("n == ", n)
	}
}
