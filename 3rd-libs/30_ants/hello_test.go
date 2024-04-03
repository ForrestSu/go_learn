package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringSliceToInts(list []string) []int {
	var ints []int
	for _, item := range list {
		if n, err := strconv.Atoi(item); err == nil {
			ints = append(ints, n)
		}
	}
	return ints
}

func Test_stringSliceToInts(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, stringSliceToInts(strings.Split("1+2+3", "+")))
	assert.Equal(t, []int{1, 2}, stringSliceToInts(strings.Split("+1+2+", "+")))
	// assert.Equal(t, []int{}, stringSliceToInts(strings.Split("", "+")))
	a := strings.Split("", "+")
	t.Logf("%d", len(a))
}
