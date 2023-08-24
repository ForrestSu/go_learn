package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

var (
	testData = []byte("hello world")
)

func TestPrint(t *testing.T) {
	raw := md5.Sum([]byte("hello world"))
	t.Logf("data=[%s]\n", hex.EncodeToString(raw[:]))
	raw2 := sha256.Sum256(testData)
	t.Logf("sha256=[%s]\n", hex.EncodeToString(raw2[:]))

}

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5.Sum(testData) // 32 char
	}
}

func BenchmarkSha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha256.Sum256(testData) // 64 char
	}
}
