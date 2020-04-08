package advance07_syncPool

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("call new...")
			return 0
		},
	}
	obj, _ := pool.Get().(int)
	t.Log(obj)
}

/**
output:
1
4
3
2
0
*/
func TestSyncPoolV2(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)

	//runtime.GC()
	for i := 0; i < 5; i++ {
		fmt.Println(p.Get())
	}
}
