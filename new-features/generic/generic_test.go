package generic

import (
	"cmp"
	"testing"
)

func TestMax(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1, 2) = %d; want 2", got)
	}
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
