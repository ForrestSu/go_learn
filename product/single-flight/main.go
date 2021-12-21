package main

import (
	"log"
	"sync"
	"time"
)

// 多等1
func testWaitGroup() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 10; i++ {
		go work(i, wg)
	}
	time.Sleep(3 * time.Second)
	log.Printf("open gate!!\n")
	wg.Done()
}

func work(id int, wg *sync.WaitGroup) {
	log.Printf("work[%d] waiting...\n", id)
	wg.Wait()
	log.Printf("work[%d] done\n", id)
}

func main() {
	testWaitGroup()
	time.Sleep(3 * time.Second)
}
