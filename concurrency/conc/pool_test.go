package main

import (
	"sort"
	"testing"

	"github.com/sourcegraph/conc/pool"
	"github.com/stretchr/testify/require"
)

func TestPool(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		g := pool.NewWithResults[int]()
		expected := []int{}
		for i := 0; i < 100; i++ {
			i := i
			expected = append(expected, i)
			g.Go(func() int {
				return i
			})
		}
		res := g.Wait()
		sort.Ints(res)
		require.Equal(t, expected, res)
	})

}
