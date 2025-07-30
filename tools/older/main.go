package main

import (
	"fmt"
)

// 养老金：公式
func main() {

	// z=0.005*(1+x)*y
	var _ = make([][]float64, 0)
	for x := 0.6; x < 3.1; x += 0.1 {
		for y := 15; y <= 30; y += 1 {
			z := 0.005 * (1 + x) * float64(y)
			fmt.Printf("[%.1f, %d, %.3f],", x, y, z)
		}
		fmt.Printf("\n")
	}
}
