package main

import (
	"testing"
	"time"
)

// 时间日期格式
const TimeFmt = "2006-01-02T15:04:05.000Z"

func TestParser(t *testing.T) {
	timeStr := "2021-01-28T08:00:00.000Z"
	tm, err := time.ParseInLocation(TimeFmt, timeStr, time.UTC)
	if err == nil {
		t.Log(tm.Local())
	} else {
		t.Error(err)
	}
}

/**
BenchmarkNow
BenchmarkNow-12                 16800339               67.8 ns/op             0 B/op           0 allocs/op
BenchmarkTimeFormat
BenchmarkTimeFormat-12           4006947               305 ns/op              32 B/op          1 allocs/op
BenchmarkParseTime
BenchmarkParseTime-12            5673578               212 ns/op               0 B/op          0 allocs/op
PASS
*/
// 获取当前时间
func BenchmarkNow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now()
	}
}

// 将time格式化为字符串
func BenchmarkTimeFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = time.Now().UTC().Format(TimeFmt)
	}
}

// 将字符串解析为time
func BenchmarkParseTime(b *testing.B) {
	b.ResetTimer()
	timeStr := "2021-01-02T15:04:05.000Z"
	for i := 0; i < b.N; i++ {
		_, _ = time.ParseInLocation(TimeFmt, timeStr, time.UTC)
	}
}
