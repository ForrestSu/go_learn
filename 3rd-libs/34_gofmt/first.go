package main

import "fmt"

// TestTask 测试任务
type testTask struct{}

// NewTestTask new
func NewTestTask() *testTask {
	return &testTask{}
}

func (t *testTask) Say() {
	fmt.Println("mock NewTestTask()")
}
