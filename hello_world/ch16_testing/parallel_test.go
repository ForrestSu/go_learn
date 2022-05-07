package ch16_testing

import "testing"

func TestParallel(t *testing.T) {
	t.Run("subtest", func(t *testing.T) {
		t.Parallel()
		if got, want := 1, 1; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
