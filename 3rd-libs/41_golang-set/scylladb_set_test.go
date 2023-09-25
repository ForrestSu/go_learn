package golang_set

import (
	"testing"

	// 非泛型库
	"github.com/scylladb/go-set/strset"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := strset.New("a", "b", "c")
	assert.True(t, s.Has("a"))
	s2 := strset.New("a")
	// Intersect set
	assert.Equal(t, []string{"a"}, strset.Intersection(s, s2).List())
}
