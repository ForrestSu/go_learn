package go1_23

import (
	"iter"
	"testing"
	"unique"
)

func TestUnique(t *testing.T) {
	words := make([]unique.Handle[string], 12)
	for i, word := range []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy"} {
		words[i] = unique.Make(word)
	}

}

func TestPull2ImmediateStop(t *testing.T) {
	next, stop := iter.Pull2(panicSeq2())
	stop()
	// Make sure we don't panic if we try to call next or stop.
	if _, _, ok := next(); ok {
		t.Fatal("next returned true after iterator was stopped")
	}
}

func panicSeq2() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		panic("boom")
	}
}
