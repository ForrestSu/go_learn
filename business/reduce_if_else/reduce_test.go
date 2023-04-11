package main

import (
	"testing"
)

// RuleFunc 根据规则，返回对应的广告位
type RuleFunc func(req *RequestMsg, rsp *ResponseMsg) int32

type Evaluate struct {
	rules []RuleFunc
}

// NewEvaluate new
func NewEvaluate(rs ...RuleFunc) Evaluate {
	return Evaluate{rules: rs}
}

// Evaluate calc
func (e Evaluate) Evaluate(req *RequestMsg, rsp *ResponseMsg) int32 {
	for _, r := range e.rules {
		if adid := r(req, rsp); adid > 0 {
			return adid
		}
	}
	return 0
}

func ConferenceAdid(req *RequestMsg, rsp *ResponseMsg) int32 {
	if rsp.IsCommercialAdFree {
		return 32
	}
	return 0
}

func TestRule(t *testing.T) {
	var req = &RequestMsg{}
	var rsp = &ResponseMsg{}
	var evaluate = NewEvaluate(ConferenceAdid)
	_ = evaluate.Evaluate(req, rsp)
}

//go:generate go test -bench=. -benchmem -cpuprofile=cpu.prof

// BenchmarkEvaluate
// BenchmarkEvaluate-12    	100000000	        11.95 ns/op
func BenchmarkEvaluate(b *testing.B) {
	var req = &RequestMsg{}
	var rsp = &ResponseMsg{}
	var evaluate = NewEvaluate(
		ConferenceAdid,
		ConferenceAdid,
		ConferenceAdid,
		ConferenceAdid,
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = evaluate.Evaluate(req, rsp)
	}
}

// BenchmarkIfElse
// BenchmarkIfElse-12    	1000000000	         0.2919 ns/op
func BenchmarkIfElse(b *testing.B) {
	var req = &RequestMsg{}
	var rsp = &ResponseMsg{}
	b.ResetTimer()
	var task = func() {
		ConferenceAdid(req, rsp)
		ConferenceAdid(req, rsp)
		ConferenceAdid(req, rsp)
		ConferenceAdid(req, rsp)
	}
	for i := 0; i < b.N; i++ {
		task()
	}
}
