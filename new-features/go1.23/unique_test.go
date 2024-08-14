package go1_23

import (
	"testing"
	"unique"
)

func TestUnique(t *testing.T) {
	words := make([]unique.Handle[string], 12)
	for i, word := range []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "fred", "plugh", "xyzzy"} {
		words[i] = unique.Make(word)
	}

}
