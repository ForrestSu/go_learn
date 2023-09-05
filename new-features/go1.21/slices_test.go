package go1_21

import (
	"slices"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/stretchr/testify/assert"
)

type Student struct {
	Name string
	Age  int
}

func TestSlices(t *testing.T) {
	var s1 = []int{1}
	var s2 = []int{1}
	assert.True(t, slices.Equal(s1, s2))

	students := []*Student{
		{Name: "Bob"},
		{Name: "Alice"},
	}
	slices.SortFunc(students, func(a, b *Student) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})
	t.Logf("%v", pretty.Sprint(students))

	slices.
}
