package jsontest

import (
	"encoding/json"
	"testing"
)

// 测试序列化
func TestSerials(t *testing.T) {

	// 将对象序列化为json对象
	e := &Employee{
		Basic: BasicInfo{
			Name: "sunquan",
			Age:  27,
		},
		Job: JobInfo{
			Skills: []string{"java", "c++", "golang"},
		},
	}
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	t.Logf("len %d, %s", len(data), data)

	// 反序列化为对象
	emp := &Employee{
		Basic: BasicInfo{},
		Job:   JobInfo{},
	}
	err1 := json.Unmarshal(data, emp)
	if err1 != nil {
		panic(err)
	}
	t.Log(emp)
}
