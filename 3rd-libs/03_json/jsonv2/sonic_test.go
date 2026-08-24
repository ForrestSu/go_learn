package jsonv2

import (
	"testing"

	"github.com/bytedance/sonic"
)

// BenchmarkSonicUnmarshal benchmarks decoding with bytedance/sonic (default config).
func BenchmarkSonicUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e Employee
		if err := sonic.Unmarshal([]byte(jsonStr), &e); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSonicMarshal benchmarks encoding with bytedance/sonic (default config).
func BenchmarkSonicMarshal(b *testing.B) {
	e := Employee{
		Basic: BasicInfo{Name: "Mike", Age: 30},
		Job:   JobInfo{Skills: []string{"Java", "Go", "C"}},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := sonic.Marshal(&e); err != nil {
			b.Fatal(err)
		}
	}
}
