package prod01_string_utils

import (
	"fmt"
	endian "github.com/virtao/GoEndian"
	"testing"
)

// 判断当前机器的大小端
func TestEndian(t *testing.T) {
	printEndian()
}

func printEndian() {
	fmt.Print("Your Machine byte order : ")
	if endian.IsBigEndian() {
		fmt.Println("Big Endian")
	} else {
		fmt.Println("Little Endian")
	}
}
