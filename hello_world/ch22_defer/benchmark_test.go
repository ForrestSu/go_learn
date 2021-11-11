package main

import (
	"sync"
	"testing"
)

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

// BenchmarkTest
// BenchmarkTest-12    	100000000	        11.27 ns/op
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

// BenchmarkTest2
// BenchmarkTest2-12    	79242217	        13.84 ns/op
func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}
