package main

import "testing"

// Test_USA_Bond
// 20年美债年收益率4.856% ,假设投入1000元,那持有20年之后到期总收益？
// year 20: 2581.46
func Test_USA_Bond(t *testing.T) {
	// 1000 * (1 + 0.04856)^20
	var base float64 = 1000
	const rate float64 = 0.04856
	const years = 20
	for i := 0; i < years; i++ {
		base = base * (1 + rate)
		t.Logf("year %d: %.2f\n", i+1, base)
	}
}
