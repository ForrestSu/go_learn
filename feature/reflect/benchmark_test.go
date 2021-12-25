package reflect

import (
	"reflect"
	"testing"
)

// 零值直接返回 false

func judgeInt64(key int64) bool {
	return key != 0
}
func judgeString(key string) bool {
	return key != ""
}

func judgeReflect(key interface{}) bool {
	return reflect.ValueOf(key).IsZero()
}

//go:generate go test -bench=. -benchmem
func BenchmarkType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeInt64(100)
	}
}
func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeString("100")
	}
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeReflect(100)
	}
}
