package jsonv2

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"testing"
)

// Employee is the data model reused across the JSON benchmark suite.
type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// JobInfo holds the employee's professional skills.
type JobInfo struct {
	Skills []string `json:"skills"`
}

// Employee is the sample payload used to benchmark JSON (de)serialization.
type Employee struct {
	Basic BasicInfo `json:"basic"`
	Job   JobInfo   `json:"job"`
}

// jsonStr mirrors the sample payload defined in the sibling json directories.
var jsonStr = `{
	"basic":{
	  	"name":"Mike",
		"age":30
	},
	"job":{
		"skills":["Java","Go","C"]
	}
}`

// BenchmarkStdUnmarshal benchmarks decoding with the classic encoding/json.
func BenchmarkStdUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e Employee
		if err := json.Unmarshal([]byte(jsonStr), &e); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkV2Unmarshal benchmarks decoding with encoding/json/v2.
func BenchmarkV2Unmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var e Employee
		if err := jsonv2.Unmarshal([]byte(jsonStr), &e); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStdMarshal benchmarks encoding with the classic encoding/json.
func BenchmarkStdMarshal(b *testing.B) {
	e := Employee{
		Basic: BasicInfo{Name: "Mike", Age: 30},
		Job:   JobInfo{Skills: []string{"Java", "Go", "C"}},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(&e); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkV2Marshal benchmarks encoding with encoding/json/v2.
func BenchmarkV2Marshal(b *testing.B) {
	e := Employee{
		Basic: BasicInfo{Name: "Mike", Age: 30},
		Job:   JobInfo{Skills: []string{"Java", "Go", "C"}},
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := jsonv2.Marshal(&e); err != nil {
			b.Fatal(err)
		}
	}
}
