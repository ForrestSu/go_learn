package db

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

//go:generate mockgen -source=db.go -destination=db_mock.go -package=db

/**
这个测试用例有2个目的:
 (1) 使用 ctrl.Finish() 断言 DB.Get() 被是否被调用，如果没有被调用，后续的 mock 就失去了意义；
 (2) 是测试方法 GetFromDB() 的逻辑是否正确(如果 DB.Get() 返回 error，那么 GetFromDB() 返回 -1)。
 */
func TestMockUsage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}

}
