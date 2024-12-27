package utils

import (
	"testing"
)

func Test_spinWheel(t *testing.T) {
	var cnt = make(map[int32]int)
	total := 100000
	for i := 0; i < total; i++ {
		kind := spinWheel(prizeList, 1)
		cnt[kind]++
	}
	// 验证概率分布
	for _, item := range prizeList {
		t.Logf("kind: %d,  count: %.3f", item.Kind, float32(cnt[item.Kind])/float32(total))
	}
}
