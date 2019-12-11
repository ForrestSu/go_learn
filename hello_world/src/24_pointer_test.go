package try_test

import "testing"

func TestPointer(t *testing.T) {
	var a int = 10
	aPtr := &a
	aPtr = aPtr + 1
	t.Log("a == ", a, ", aPtr == ", aPtr)
	t.Logf("%T %T", a, aPtr)
}