package main

import (
	"fmt"
	"testing"

	"github.com/scylladb/go-set/strset"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func Print[T ~int | ~string](t T) {
	fmt.Printf("%v", t)
}

func TestExp(t *testing.T) {
	var s1 = []string{"1", "2"}
	var s2 = []string{"1", "2"}
	assert.True(t, slices.Equal(s1, s2))
	strset.New()
	Print(100)
}
