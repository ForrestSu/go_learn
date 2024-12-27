package random_bench

import (
	"math/rand"
	randv2 "math/rand/v2"
	"testing"
)

func BenchmarkRandV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randv2.IntN(1000)
	}
}

func BenchmarkRandV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}
