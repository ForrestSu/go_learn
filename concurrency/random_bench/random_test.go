package random_bench

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// BenchmarkRand-10    	  102932    10928 ns/op
func BenchmarkRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var rndGen = rand.New(rand.NewSource(1))
		rndGen.Intn(10)
	}
}

func Benchmark_GoroutineSafe(b *testing.B) {
	var rndGen = rand.New(rand.NewSource(1))
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		rndGen.Intn(10)
		mu.Unlock()
	}
}

// BenchmarkRandReuse-10 	304261048   3.892 ns/op
func BenchmarkRandReuse(b *testing.B) {
	var rndGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		rndGen.Intn(10)
	}
}

func TestOk(t *testing.T) {
	t.Log(rand64())
}

func BenchmarkInt63ThreadsafeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(10)
		}
	})
}
