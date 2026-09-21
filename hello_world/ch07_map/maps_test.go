package try

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaps(t *testing.T) {
	var m = map[string]string{"a": "1", "b": "2", "c": "3"}
	// collect and sort
	sortedKeys := slices.Sorted(maps.Keys(m))
	assert.Equal(t, []string{"a", "b", "c"}, sortedKeys)
}
