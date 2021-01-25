package main

import "fmt"

const (
	InvalidParamVid = iota
	InvalidParam2
)

// 匿名函数只能用来做语法检查，无法被显式调用 (通过reflect也不行)
func _() {
	var x [1]struct{}
	_ = x[InvalidParamVid-0]
	_ = x[InvalidParam2-1]
}

func init() {
	fmt.Println("init...")
}
