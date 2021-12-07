package ch09_func_test

import (
	"fmt"
	"testing"
)

func change(ref ...string) {
	ref[0] = "Go"
	// ref 是一个新的切片对象，
	// ref 和 welcome 引用相同的数组
	fmt.Printf("ref address: %p\n", &ref)
	ref = append(ref, "playground")
	fmt.Println(ref)
	fmt.Printf("ref address: %p\n", &ref)
}

func TestVariadic(t *testing.T) {
	str := []string{"hello", "world"}
	fmt.Printf("str address: %p, data %s.\n", &str, str)
	change(str...)
	fmt.Printf("after str address: %p, data %s.\n", &str, str)
}

/**
output:
str address: 0xc000004480
ref address: 0xc0000044e0
[Go world playground]
ref address: 0xc0000044e0
after str address: 0xc000004480, data [Go world].
*/
