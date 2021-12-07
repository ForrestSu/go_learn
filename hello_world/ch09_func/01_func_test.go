package ch09_func

import (
	"fmt"
	"testing"
	"time"
)

// 入参 和 返回值 可以省略
func TestFunc(t *testing.T) {
	t.Log("this is a sample func.")
}

func returnMultiResult(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// 01 测试多个返回值
func TestFn(t *testing.T) {
	area, _ := returnMultiResult(10, 20)
	t.Log("area: ", area)
}

// 可以预先定义返回值，
// return 不用显示指定，在遇到return语句时会自动从函数返回它们。
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

// 函数式编程：把传入的函数inner 包装一层，返回一个新的函数
func counterSpentTime(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		result := inner(n)
		fmt.Println("cost seconds: ", time.Since(start))
		return result
	}
}

// 定义一个耗时函数
func spentFun(op int) int {
	time.Sleep(5 * time.Second)
	return op * op
}

// 02 函数是一等公民,采用函数式编程，测试耗时
func TestCost(t *testing.T) {
	newFN := counterSpentTime(spentFun)
	t.Log(newFN(10))
}
