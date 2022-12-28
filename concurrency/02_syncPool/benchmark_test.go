package syncpool

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var buf, _ = json.Marshal(Student{Name: "Forrest", Age: 25})

// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkUnmarshal
// BenchmarkUnmarshal-12                      13760             87107 ns/op            1400 B/op          8 allocs/op
// BenchmarkUnmarshalWithPool
// BenchmarkUnmarshalWithPool-12              13818             86874 ns/op             248 B/op          7 allocs/op
func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

//go:generate go test -bench . -benchmem
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
