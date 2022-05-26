package prod01_string_utils

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var nativeEndian binary.ByteOrder

func init() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)
	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
		fmt.Println(">>> LittleEndian")
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
		fmt.Println("BigEndian")
	default:
		panic("Could not determine native endianness.")
	}
}
