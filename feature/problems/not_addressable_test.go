package problems

import (
	"fmt"
	"testing"
)

type ByteSlice int32

func (b ByteSlice) valueFunc() {
	fmt.Println("value func!")
}

func (b *ByteSlice) pointerFunc() {
	fmt.Println("pointer func!")
}

func TestValueCall(t *testing.T) {
	const b ByteSlice = 123
	// t.Logf("%T", b)
	b.valueFunc() //ok
	// 不可寻址的值对象，不能调用指针方法
	b.pointerFunc() //error
}
