package main_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoto(t *testing.T) {
	var a = 0
LOOP:
	for ; a < 10; a++ {
		// skip print 5
		if a == 5 {
			a++
			goto LOOP
		}
		fmt.Printf("a = : %d\n", a)
	}
}

func TestGoto2(t *testing.T) {

	arr := strings.Split("3+4", "+")
	assert.EqualValues(t, []string{"3", "4"}, arr)
	t.Logf("<%v>, len=%d\n", arr, len(arr))

	arr = strings.Split("3", "+")
	assert.EqualValues(t, []string{"3"}, arr)
	t.Logf("<%v>, len=%d\n", arr, len(arr))

	arr = strings.Split("3+", "+")
	assert.EqualValues(t, []string{"3", ""}, arr)
	t.Logf("<%v>, len=%d\n", arr, len(arr))

	arr = strings.Split("", ",")
	assert.EqualValues(t, []string{""}, arr)
	t.Logf("<%v>, len=%d\n", arr, len(arr))
}
