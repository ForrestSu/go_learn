package _2_lodash

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestLO(t *testing.T) {
	var arr = []int{1, 2, 3}
	// reverse slice (original slice will be changed!)
	assert.Equal(t, []int{3, 2, 1}, lo.Reverse(arr))
	t.Log(arr)
	// count occurrences of 1
	assert.Equal(t, 1, lo.Count(arr, 1))
}

func Reverse[T any](slice []T) []T {
	length := len(slice)
	half := length / 2
	for i := 0; i < half; i++ {
		j := length - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
