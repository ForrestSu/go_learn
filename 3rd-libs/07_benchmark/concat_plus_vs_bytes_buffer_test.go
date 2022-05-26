package benchmark

import (
	"bytes"
	"testing"
)

/**
go test -bench=. -v -benchmem
*/

func BenchmarkStringConcatByAdd(b *testing.B) {
	//b.Log(b.N)
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuffer(b *testing.B) {
	//b.Log(b.N)
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
		//_ = buf.String()
	}

	b.StopTimer()
}
