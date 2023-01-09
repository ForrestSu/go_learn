package conc

import (
	"sort"
	"testing"

	"github.com/sourcegraph/conc/pool"
	"github.com/stretchr/testify/require"
)

func TestPool(t *testing.T) {
	// process
	// iter.ForEach(nil, handle)
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

func process(stream chan int) {
	p := pool.New().WithMaxGoroutines(10)
	for elem := range stream {
		elem := elem
		p.Go(func() {
			handle(elem)
		})
	}
	p.Wait()
}

func handle(v int) {}
