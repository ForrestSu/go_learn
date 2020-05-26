package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestArguments(t *testing.T) {
	fmt.Println("print:")
	// 打印命令行参数
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello", os.Args[1])
	}
}
