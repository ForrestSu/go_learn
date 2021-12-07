package try_test

import (
	"testing"
)

// map 值对象可以是方法
// 方法在 Go 里面是一等公民
func TestSetByMap(t *testing.T) {
	mySet := map[int]bool{}
	// 1 添加元素
	mySet[10] = true
	mySet[20] = true

	// 4 元素个数
	t.Log(len(mySet))
	// 2 判断一个元素是否存在
	if mySet[10] {
		t.Log("existed!")
	} else {
		t.Log("not existed!")
	}
	// 3 删除元素
	delete(mySet, 10)

	// 再次检查这个 元素是否存在
	if mySet[10] {
		t.Log("existed!")
	} else {
		t.Log("not existed!")
	}
}
