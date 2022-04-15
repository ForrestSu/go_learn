package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func slowJob(id int) string {
	log.Printf("slow job %v running...\n", id)
	time.Sleep(3 * time.Second)
	return fmt.Sprintf("return by job %v!", id)
}

// lambda pass object by reference
func runTest() {
	g := Group{}
	log.Printf("address(p) = %p", &g)
	// simulate multiple jobs request
	for i := 0; i < 5; i++ {
		go func(id int) {
			result, _ := g.Do("userid-123", func() (interface{}, error) {
				return slowJob(id), nil
			})
			log.Printf("address(p) = %p, job %v => %v\n", &g, id, result)
		}(i)
	}
	time.Sleep(5 * time.Second)
}

func TestGroup_Do(t *testing.T) {
	runTest()
}
