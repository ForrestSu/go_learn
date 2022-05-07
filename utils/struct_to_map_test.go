package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want map[string]string
	}{
		{name: "string",
			in: &struct {
				Name string `json:"name"`
			}{Name: "John"},
			want: map[string]string{"name": "John"},
		},
		{name: "age",
			in: struct {
				Age int `json:"age"`
			}{Age: 10},
			want: map[string]string{"age": "10"},
		},
		{name: "slice",
			in: struct {
				Slice []string `json:"slice"`
			}{Slice: []string{"a", "b", "c"}},
			want: map[string]string{"slice": "[\"a\",\"b\",\"c\"]"},
		},
		{name: "nil slice",
			in: struct {
				NilSlice []string `json:"nil_slice"`
			}{},
			want: map[string]string{"nil_slice": "[]"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, StructToMap(tt.in), "StructToMap(%v)", tt.in)
		})
	}
}
