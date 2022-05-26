package diff

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"

	"github.com/google/go-cmp/cmp"
)

type Student struct {
	Name  string
	Age   int
	Tel   string
	Class string
}

func TestEqual(t *testing.T) {
	want := &Student{
		Name: "Bob",
		Age:  20,
	}
	got := &Student{
		Name: "Bob",
		Age:  21,
	}
	// 比较时, 忽略部分字段
	var opt = cmpopts.IgnoreFields(Student{}, "Age")
	assert.True(t, cmp.Equal(want, got, opt))

	// 自定义差异记录器
	var r DiffReporter
	if !cmp.Equal(want, got, cmp.Reporter(&r)) {
		t.Logf("simple diff (-want +got):\n%s", r.String())
	}
}

func TestDiff(t *testing.T) {
	want := &Student{
		Name: "Bob",
		Age:  20,
	}
	got := &Student{
		Name: "Bob",
		Age:  21,
	}
	diff := cmp.Diff(want, got)
	t.Logf("Student{} mismatch (-want +got):\n%s", diff)
}
