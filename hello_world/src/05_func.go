package main

import "fmt"

// 入参 和 返回值 可以省略
func testFunc() {
	fmt.Println("this is a sample func.")
}

func testFuncParam(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}


// 可以预先定义返回值，
// return 不用显示指定，在遇到return语句时会自动从函数返回它们。
func rectProps(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}


func main() {
	testFunc()
	// 丢弃参数
	area, _ := testFuncParam(10, 20)
	fmt.Println("area: ", area )
}
