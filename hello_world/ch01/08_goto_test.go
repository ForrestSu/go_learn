package main_test

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	var a = 0
LOOP:
	for ;a < 10;a++ {
		/* skip print 5 */
		if a == 5 {
			a++
			goto LOOP
		}
		fmt.Printf("a = : %d\n", a)
	}
}
