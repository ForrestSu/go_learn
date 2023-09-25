package golang_set

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenericsSet(t *testing.T) {
	s := mapset.NewThreadUnsafeSet[string]("a", "b", "c")
	assert.True(t, s.Contains("a"))
	s2 := mapset.NewThreadUnsafeSet[string]("a")
	// Intersect set
	assert.Equal(t, []string{"a"}, s.Intersect(s2).ToSlice())
}
