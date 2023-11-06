package main

import (
	"fmt"
	"runtime"

	// docker 容器正确获取cpu核数
	_ "go.uber.org/automaxprocs"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
}
