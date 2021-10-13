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
		Name: "zhangsan",
		Age:  20,
	}
	got := &Student{
		Name: "zhangsan",
		Age:  21,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
	} else {
		t.Log("no difference.")
	}
}
