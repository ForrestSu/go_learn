package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	log.Println("start...")
	for cnt := 0; cnt < 3; cnt++ {
		// wait
		<-ticker.C
		log.Printf("count %d!", cnt)
	}
	fmt.Println("done")
}
