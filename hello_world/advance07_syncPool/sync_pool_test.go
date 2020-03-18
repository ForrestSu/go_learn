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
