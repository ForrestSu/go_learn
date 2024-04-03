package main

import "testing"

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Invoke(doWork)
		// doWork()
		compare()
	}
}

// func BenchmarkPlain(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		doWork()
//	}
// }

func doWork() {
	for i := 0; i < 1; i++ {
	}
}
func compare() {
	for i := 0; i < 1; i++ {
	}
}

func Invoke(v interface{}) {
	if fn, ok := v.(func()); ok {
		fn()
	}
}
