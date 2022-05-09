package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Student struct {
	Name string
	Age  int
}

func TestCmp(t *testing.T) {
	want := &Student{
		Name: "Bob",
		Age:  20,
	}
	got := &Student{
		Name: "Bob",
		Age:  21,
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Logf("Student{} mismatch (-want +got):\n%s", diff)
	}
}
