package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
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

func TestMockNot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Tom")).Return(0, nil)

	var val = GetFromDB(m, "ZhangSan")
	assert.Equal(t, 0, val)
}

/**
通过InOrder关键字实现的保序示例代码：

gomock.InOrder(
	mockRepo.EXPECT().Retrieve(Any()).Return(nil, ErrAny)
	mockRepo.EXPECT().Create(Any(), Any()).Return(nil)
	mockRepo.EXPECT().Retrieve(Any()).Return(objBytes, nil)
)

当mock对象行为的注入保序后，如果行为调用的顺序和其不一致，就会触发测试失败。
如果在测试用例执行过程中，
Repository方法的调用顺序如果不是按 Retrieve -> Create -> Retrieve的顺序进行，则会导致测试失败。
*/
