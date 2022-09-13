package main

import (
	"context"
	"testing"
)

func Test_runTask(t *testing.T) {
	type args struct {
		ctx   context.Context
		secTO int
		n     *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTask(tt.args.ctx, tt.args.secTO, tt.args.n)
		})
	}
}
