package utils

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var info = struct {
	Age        int      `json:"age"`
	EmptySlice []string `json:"empty_slice"`
	Name       string   `json:"name"`
	Slice      []string `json:"slice"`
}{
	Name:  "test",
	Age:   10,
	Slice: []string{"a", "b", "c"},
}

func TestStructToMap(t *testing.T) {
	val := StructToMap(&info)
	data, err := json.Marshal(val)
	assert.Nil(t, err)
	expect := `{"age":"10","empty_slice":"[]","name":"test","slice":"[\"a\",\"b\",\"c\"]"}`
	assert.Equal(t, expect, string(data))
}

// BenchmarkStructToMap
// BenchmarkStructToMap-12 1396696  859.0 ns/op   424 B/op   9 allocs/op
func BenchmarkStructToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StructToMap(&info)
	}
}
