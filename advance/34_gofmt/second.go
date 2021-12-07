package main

import "fmt"

// SecondTask 测试任务
type SecondTask struct{}

// NewSecondTask new
func NewSecondTask() *SecondTask {
	return &SecondTask{}
}

func (t *SecondTask) Say() {
	fmt.Println("mock NewSecondTask()")
}
