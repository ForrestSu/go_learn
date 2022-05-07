package ch16_testing

import (
	"log"
	"runtime"
	"testing"
	"time"
)

//go:generate go test -parallel 5 -v .

// 并行表示该测试将与（并且仅与）其他并行测试并行运行。
// 当由于使用 -test.count 或 -test.cpu 多次运行测试时，单个测试的多个实例永远不会彼此并行运行。
func TestA1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arg, want string
	}{
		{"A0", "A0"},
		{"A1", "A1"},
		{"A2", "A2"},
		{"A3", "A3"},
	}
	for i, test := range tests {
		var echoMsg = func(msg string) string {
			time.Sleep(1 * time.Second)
			log.Printf("case %d, count(g): %d, %s\n", i, runtime.NumGoroutine(), msg)
			return msg
		}
		if got := echoMsg(test.arg); got != test.want {
			t.Errorf("echoMsg(%q) = %q, want %q", test.arg, got, test.want)
		}
	}
}

func TestC2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arg, want string
	}{
		{"P0", "P0"},
		{"P1", "P1"},
		{"P2", "P2"},
		{"P3", "P3"},
	}
	for i, test := range tests {
		t.Run("sub_test", func(t *testing.T) {
			t.Parallel()
			var echoMsg = func(msg string) string {
				time.Sleep(1 * time.Second)
				log.Printf("==> %d, %s\n", i, msg)
				return msg
			}
			if got := echoMsg(test.arg); got != test.want {
				t.Errorf("echoMsg(%q) = %q, want %q", test.arg, got, test.want)
			}
		})
	}
}
