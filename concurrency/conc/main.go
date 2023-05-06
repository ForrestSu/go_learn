package main

import (
	"fmt"

	"github.com/sourcegraph/conc/pool"
)

func main() {
	ExamplePool()
}

func ExamplePool() {
	p := pool.New().WithMaxGoroutines(3)
	for i := 0; i < 5; i++ {
		k := i
		p.Go(func() {
			fmt.Println("conc=", k)
		})
	}
	p.Wait()
	// Output:
	// conc
	// conc
	// conc
	// conc
	// conc
}
