package reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// 零值直接返回 false

func judgeInt64(key int64) bool {
	return key != 0
}
func judgeString(key string) bool {
	return key != ""
}

var data = map[string]bool{
	"100": true,
	"102": true,
}

func init() {
	for i := 0; i < 10000; i++ {
		data[strconv.Itoa(i)] = true
	}
	fmt.Println("inited...")
}

func judgeMap(key string) bool {
	return data[key]
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

func BenchmarkMapFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeMap("100")
	}
}

func BenchmarkReflectIsZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeReflect(100)
	}
}

func BenchmarkStrToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.Atoi("100")
	}
}
