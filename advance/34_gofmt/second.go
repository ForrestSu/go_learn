package main

import "fmt"

// SecondTask 测试任务
type secondTask struct{}

// NewSecondTask new
func NewSecondTask() *secondTask {
	return &secondTask{}
}

func (t *secondTask) Say() {
	fmt.Println("mock NewSecondTask()")
}
