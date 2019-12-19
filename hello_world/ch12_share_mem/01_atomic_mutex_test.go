package ch12_share_mem

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 在不同的协程自增，需要对共享资源进行互斥访问
// counter = 4718
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter = %d\n", counter)
}

// 1 使用 sync.Mutex 保护共享变量
func TestMutex(t *testing.T) {
	counter := 0
	mtx := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		go func() {
			mtx.Lock()
			defer mtx.Unlock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter = %d\n", counter)
}

// 2 使用原子操作
func TestAtomic(t *testing.T) {
	var counter int32 = 0
	for i := 0; i < 5000; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
		}()
	}
	time.Sleep(time.Second)
	t.Logf("counter = %d\n", counter)
}

// 针对 读操作比较多的情况，使用 rwlock

func TestRwlock(t *testing.T) {
	//mtx := sync.RWMutex{}
}
