package main

import (
	"fmt"
	"os"
	"runtime"

	// docker 容器正确获取cpu核数
	_ "go.uber.org/automaxprocs"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(runtime.GOMAXPROCS(0))
}
