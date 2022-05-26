package main

import (
	"fmt"
	"strconv"
	"testing"
)

/**
 * FormatInt more faster!
 * pkg: github.com/ForrestSu/go_learn/feature/prod06_perf
 * BenchmarkSprintf-4      11784390                95.7 ns/op            16 B/op          1 allocs/op
 * BenchmarkFormatInt-4    38684365                32.1 ns/op             7 B/op          0 allocs/op
 * PASS
 */

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", int64(i))
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(int64(i), 10)
	}
}
