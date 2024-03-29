package proto

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 安装 easyjson； 用法：easyjson -all <file>.go
//go:generate go install  github.com/mailru/easyjson/easyjson

var jsonStr = `{
	"basic":{
	  	"name":"Mike",
		"age":30
	},
	"job":{
		"skills":["Java","Go","C"]
	}
}`

func TestEasyJSON(t *testing.T) {
	emp := &Employee{}
	err := emp.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Log(emp)
	if str, err := emp.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(str))
	}
}

// benchmark

func BenchmarkEmbeddedJson(b *testing.B) {
	// b.ResetTimer()
	b.StartTimer()
	e := new(Employee)
	// b.Log(b.N)
	for i := 0; i < b.N; i++ {

		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

func BenchmarkEasyJson(b *testing.B) {
	e := Employee{}
	b.Log(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
