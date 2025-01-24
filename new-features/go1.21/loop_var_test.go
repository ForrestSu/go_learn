package go1_21

import (
	"testing"
	"time"
)

// 在go1.22之前，循环变量是共享的。（版本以 go.mod 配置的版本为准）
// （即使没有并发，也会有共享问题）
func Test_LoopVar(t *testing.T) {
	taskMap := map[string][]string{
		"k1": {"x", "1"},
		"k2": {"y", "2"},
		"k3": {"z", "3"},
	}
	var fns []func()
	for _, v := range taskMap {
		copyVal := v // avoid share loop variables in closure
		fns = append(fns, func() {
			// t.Logf("k=%s, address: (&k)=%p, (&copyK)=%p, ", k, &k, &copyK)
			t.Logf("address: (&val)=%p, (&copyK)=%p", &v, &copyVal)
			t.Logf("v=%s, copyVal=%v", v, copyVal)
		})
	}
	for _, fn := range fns {
		fn()
	}
	time.Sleep(1 * time.Second)
}
