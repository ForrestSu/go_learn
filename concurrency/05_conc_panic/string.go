package main

import (
	"fmt"
	"time"
)

func main() {
	raceString()
}

func raceString() {
	rwStr := "init"
	// read only
	go func() {
		for {
			time.Sleep(10 * time.Nanosecond)
			fmt.Sprintf("str = %s", rwStr)
		}
	}()
	// write only
	for {
		rwStr = ""
		time.Sleep(10 * time.Nanosecond)
		rwStr = "/test/test/test"
		time.Sleep(10 * time.Nanosecond)
	}
}
