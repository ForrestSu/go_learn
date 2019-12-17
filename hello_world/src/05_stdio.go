package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//以换行符结束
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}