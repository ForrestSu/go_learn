package try_test

import (
	"testing"
)

// Go默认会做初始化
func TestArrayInit(t *testing.T) {
	//Go 语言会默认将数组元素初始化为0
	var arr [3]int
	t.Log(arr[0], arr[1])

	arr2 := [4]int{1, 2, 3, 4}
	t.Log(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr3)
}

//数组遍历
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for index, ele := range arr3 {
		t.Logf("arr[%d] = %d", index, ele)
	}
	// Go 语言是个强约束的语言，声明变量必须使用
	for _, ele := range arr3 {
		t.Log(ele)
	}
}

//Go 数组截取
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3Sec := arr3[3:]
	t.Log(arr3Sec) // [4,5]
}
