package ch16_testing

import "testing"

func TestOK(t *testing.T) {
	t.Log("begin...")
	t.Log("ok!")
	t.Log("end.")
}

func TestAbnormalCase1(t *testing.T) {
	t.Log("begin...")
	t.Error("error occur!")
	t.Log("end.")
}

func TestAbnormalCase2(t *testing.T) {
	t.Log("begin...")
	t.Fatal("error occur!")
	//unreachable here...
	t.Log("end.")
}

func TestOK2(t *testing.T) {
	t.Log("begin...")
	defer t.Log("end.")
	// Fatal will run Goexit()
	t.Fatal("error occur!")
}
