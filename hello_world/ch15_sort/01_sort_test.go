package ch15_sort

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	arrSlice := []int{5, 1, 2, 3, 4}
	sort.Ints(arrSlice)
	t.Log(arrSlice)

	str := []string{"55", "44", "3", "2", "1"}
	sort.Strings(str)
	t.Log(str)
}

// 自定义类型排序
type Student struct {
	id    int
	name  string
	score int
}

type StudentSlice []Student

func (a StudentSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a StudentSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a StudentSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].score > a[i].score
}
func TestSortStruct(t *testing.T) {
	stu := []Student{
		{1, "zhang san", 52},
		{2, "li si", 30},
		{3, "wang wu", 12},
		{4, "zhao liu", 26},
	}
	sort.Sort(StudentSlice(stu))
	t.Log(stu)
}
