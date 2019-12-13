package ch09_func

import (
	"fmt"
	"testing"
)

/**
output:
print 0
print 3
print 2
print 1
 */
func TestDefer(t *testing.T) {

	fmt.Println("print 0")
	// push stack
	defer fmt.Println("print 1")
	// push stack
	defer fmt.Println("print 2")

	fmt.Println("print 3")
}

func TestDeferFunc(t *testing.T){
	defer func(){
		t.Log("Clean resources.")
	}()
	t.Log("Started")
	panic("Fatal error") //defer 仍会执行
	//t.Log("aa") //unreachable code
}

