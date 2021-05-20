package fastid

import (
	"testing"
)

func TestGetMachineByIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want byte
	}{
		{name: "case_empty", ip: "", want: 0},
		{name: "case_ok", ip: "172.0.210.1", want: 210 ^ 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineByIP(tt.ip); got != tt.want {
				t.Errorf("GetMachineByIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
