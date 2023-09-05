package main

import (
	"fmt"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	pool, err := ants.NewPoolWithFunc(1000, func(arg interface{}) {
		if fn, ok := arg.(func()); ok {
			fn()
		}
	})
	if err != nil {
		return
	}
	err = pool.Invoke(func() {
		fmt.Println("invoke")
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
}
