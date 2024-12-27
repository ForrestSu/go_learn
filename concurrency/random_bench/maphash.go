package random_bench

import (
	"hash/maphash"
	_ "unsafe"
)

//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64

func randInt64() uint64 {
	return runtime_rand()
}

// 生成
func rand64() uint64 {
	return maphash.Bytes(maphash.MakeSeed(), nil)
}
