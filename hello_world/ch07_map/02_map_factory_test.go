package try_test

import (
	"testing"
)

// map 值对象可以是方法
// 方法在 Go 里面是一等公民
func TestFactory(t *testing.T) {
	factory := map[int]func(int) int{}
	factory[1]= func (op int) int {return op}
	factory[2]= func (op int) int {return op*op}
	factory[3]= func (op int) int {return op*op*op}
	// func(int) int
	t.Logf("%T", factory[1])
	//map[1:0x10f8c80]
	t.Log(factory)

	n := 10
	for i:=1; i<=3;i++{
		value := factory[i](n)
		t.Log(n, value)
	}
}
