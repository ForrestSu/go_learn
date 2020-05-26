package main_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestScanf(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	//以换行符结束
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}
