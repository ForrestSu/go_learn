package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	// IpcCreate create if key is nonexistent
	IpcCreate = 00001000
)

var mode = flag.Int("mode", 0, "0:write 1:read")

//go:generate ./dev02_syscall -mode 1

func main() {
	flag.Parse()
	key := uintptr(0x02)
	memSize := uintptr(4)
	shmId, _, err := syscall.Syscall(syscall.SYS_SHMGET, key, memSize, IpcCreate|0600)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("shmId: %v\n", shmId)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmId, 0, 0)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-2)
	}
	fmt.Printf("shmaddr: %v\n", shmaddr)

	defer syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)

	if *mode == 0 {
		fmt.Println("write mode")
		i := 0
		for {
			fmt.Printf("%d\n", i)
			*(*int)(unsafe.Pointer(uintptr(shmaddr))) = i
			i++
			time.Sleep(1 * time.Second)
		}
	} else {
		fmt.Println("read mode")
		for {
			fmt.Println(*(*int)(unsafe.Pointer(uintptr(shmaddr))))
			time.Sleep(1 * time.Second)
		}
	}
}
