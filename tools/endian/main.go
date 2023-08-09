package main

import (
	"unsafe"
)

func IsBigEndian() bool {
	var t uint32 = 0x12345678
	var pb = (*byte)(unsafe.Pointer(&t))
	return *pb == 0x12
}

func main() {
	if IsBigEndian() {
		println("大端序[BigEndian]")
	} else {
		println("小端序")
	}
}
