package prod10_json_bench

import (
	"encoding/json"
	"testing"
)

type Point struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

// BenchmarkJson
// BenchmarkJson-12    	 1848033	       622.5 ns/op
func BenchmarkJson(b *testing.B) {
	var text = []byte(`{"x":10,"y":20}`)
	var point = &Point{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(text, point)
	}
}
