package main

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// TODO: 如果为空的 struct，则每次new都返回相同的对象
type SingleObj struct {
	name string
}

var obj *SingleObj
var once sync.Once

// once 内部采用 mutex 实现
func singleton() *SingleObj {
	once.Do(func() {
		fmt.Println("Create Singleton")
		obj = &SingleObj{name: "hah"}
		fmt.Printf("==> %T, %x\n", obj, unsafe.Pointer(obj))
	})
	return obj
}

func TestSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := singleton()
			fmt.Printf("%T, %x\n", obj, unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
