package ch17_obj_pool

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	pool := NewObjPool(5)
	t.Logf("new: %+v", pool)

	// 使用对象
	conn0, _ := pool.Take(time.Second)
	for i := 0; i < 6; i++ {
		if conn, err := pool.Take(time.Second); err != nil {
			t.Log("err:", err)
		} else {
			t.Logf("%+v", conn)
		}
	}

	// 归还
	t.Log("give back 1 conn!")
	if err := pool.Release(conn0); err != nil {
		t.Fatal("fail to release 1 conn")
	}
	if conn, err := pool.Take(time.Second); err != nil {
		t.Log("err:", err)
	} else {
		t.Logf("last %+v", conn)
	}
	t.Log("ok!")
}
