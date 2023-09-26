package main

import "fmt"

func main() {
	p := NewPoolWithSize(10)
	p.Start()
	for i := 0; i < 3; i++ {
		p.AddJob(func(id int) {
			fmt.Printf("worker %d, hello\n", id)
		})
	}
	p.ExitWait()
	p.Stop()
}
