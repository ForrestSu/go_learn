package utils

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadFileToLines(t *testing.T) {
	var text = "hello\n\n"
	type args struct {
		r           io.Reader
		ignoreEmpty bool
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "case_ok",
			args: args{r: strings.NewReader(text), ignoreEmpty: true},
			want: []string{strings.TrimSpace(text)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileToLines(tt.args.r, tt.args.ignoreEmpty)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileToLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileToLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
