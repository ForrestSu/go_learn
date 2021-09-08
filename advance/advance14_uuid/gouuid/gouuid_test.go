package gouuid

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUuid(t *testing.T) {
	v1 := uuid.NewV1()
	t.Log(v1.String(), v1.Version())

	// uuid.FromString()
}

// BenchmarkUUID1
// BenchmarkUUID1-12    	 7106551	       156.2 ns/op
// BenchmarkUUID4
// BenchmarkUUID4-12    	 1515741	       760.0 ns/op
func BenchmarkUUID1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.NewV1().String()
	}
}

func BenchmarkUUID4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.NewV4().String()
	}
}
