package advance14_uuid

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestParseUUID(t *testing.T) {
	id, _ := uuid.NewUUID()
	strUUID := id.String()
	t.Log(id.Version(), strUUID)

	// VERSION_4 304186278164081700 Domain40
	// "b99a75c8-36a7-442e-8a34-08b55ab49e32"  time 301238319081747912
	githubUUID, _ := uuid.Parse("5fd24380-a81f-11eb-93ad-7d82ade035ea")
	t.Log(githubUUID.Version(),
		time.Unix(githubUUID.Time().UnixTime()),
		githubUUID.Domain(),
	)
}

// BenchmarkUUID1
// BenchmarkUUID1-12    	 7487606	       160.5 ns/op
// BenchmarkUUID4
// BenchmarkUUID4-12    	 1539740	       788.2 ns/op
func BenchmarkUUID1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v1, _ := uuid.NewUUID()
		_ = v1.String()
	}
}
func BenchmarkUUID4(b *testing.B) {
	uuid.EnableRandPool()
	for i := 0; i < b.N; i++ {
		_ = uuid.New().String()
		uuid.NewString()
	}
}
