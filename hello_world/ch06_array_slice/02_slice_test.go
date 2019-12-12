package try_test

import (
	"testing"
)

//slice cap 大小为倍增
func TestSlice(t *testing.T) {
	var s0 []int
	t.Log("len(s0) == ", len(s0), cap(s0))

	// 在切片后追加元素
	for n := 1; n <= 3; n++ {
		s0 = append(s0, n)
		t.Log("len(s0) == ", len(s0), cap(s0))
	}
	// 02 定义一个切片
	s1 := []int{1, 2, 3}
	t.Log("len(s1) == ", len(s1), cap(s1))

	//03 定义一个长度为 3, 容量为 5 的切片
	s2 := make([]int, 3, 5)
	t.Log("len(s2) = ", len(s2), "cap(s2) = ", cap(s2))
	t.Log(s2)
}

func TestSliceGrowing(t *testing.T) {
	var s0 []int
	t.Log("len(s0) == ", len(s0), cap(s0))
	// 在切片后追加元素
	for n := 1; n <= 10; n++ {
		s0 = append(s0, n)
		t.Logf("%p, len(s0) == %d,%d", &s0, len(s0), cap(s0))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := [...]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul",
		"Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	// [Apr May Jun] 3 9
	t.Log(Q2, len(Q2),cap(Q2))

	// [Jun Jul Aug] 3 7
	summer := year[5:8]
	t.Log(summer, len(summer),cap(summer))

	//"June" => "Unknown"
	summer[0] = "Unknown"
	t.Log(year)
}

func TestSliceCompare(t *testing.T){
	a := []int {1,2,3}
	b := a[2:]
	// invalid operation: a == b (slice can only be compared to nil)
	if(a == b){

	}
}