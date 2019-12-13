package ch09_func

import (
	"testing"
)

// 变长参数，求和
func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func TestVariadicParams(t *testing.T) {
	//sum(1..5) = 15
	t.Log(Sum(1,2,3,4,5))
}
