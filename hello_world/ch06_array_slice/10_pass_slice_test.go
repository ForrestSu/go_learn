package try_test

import (
	"fmt"
	"testing"
)

// 切片和python 切片类似，注意切片是引用，可以使用copy 进行深拷贝
func TestArray(t *testing.T) {
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a)
	fmt.Println("length of a is", len(a))
}

// 数组，按值传递，深Copy一份
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func TestPassByValue(t *testing.T) {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

// 由于切片是共享存储，切片对象这里会copy一份，但是堆上的元素数组是shared的.
func changeSlice(num []int) {
	num[0] *= 10
	num = append(num, 4)
	fmt.Println("slice inside function ", num)
}

// 采用指针，传入的切片对象会被修改
func passSlicePointer(num *[]int) {
	(*num)[0] *= 10
	*num = append(*num, 4)
	fmt.Println("slice inside function ", num)
}

func TestSlicePassReference(t *testing.T) {
	a := []int{1, 2, 3} // ... makes the compiler determine the length
	fmt.Println(a)
	changeSlice(a)
	fmt.Println("after:", a)
	fmt.Println("length of a is", len(a))
	//采用指针传参, 能够完整带出 slice 对象属性
	passSlicePointer(&a)
	fmt.Println("pointer:", a)
	fmt.Println("length of a is", len(a))
}
