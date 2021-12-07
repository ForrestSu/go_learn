package main

import "fmt"

// TestTask 测试任务
type TestTask struct {
}

// NewTestTask new
func NewTestTask() *TestTask {
	return &TestTask{}
}

func (t *TestTask) Say() {
	fmt.Println("mock NewTestTask()")
}
