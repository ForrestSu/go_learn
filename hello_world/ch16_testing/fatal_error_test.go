package ch16_testing

import "testing"

func TestOK(t *testing.T) {
	t.Log("hello world")
}

/*
t.Error() vs t.Fatal()
相同点：
  都用于标记-当前单测失败
不同点：
  t.Error() 会继续执行当前单测的后续代码
  t.Fatal() 会中止当前测试用例，执行下一个case
*/

/*
func TestAbnormalCase1(t *testing.T) {
	t.Log("start...")
	t.Error("error occur!")
	t.Log("will execute....")
}

func TestAbnormalCase2(t *testing.T) {
	t.Log("begin...")
	t.Fatal("error occur!")
	// unreachable here...
	t.Log("not execute here...")
}

func TestFailure(t *testing.T) {
	t.Log("begin...")
	defer t.Log("end.")
	// Fatal will run Goexit()
	t.Fatal("error occur!")
}
*/
