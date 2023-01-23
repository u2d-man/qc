package cmd

import (
	"testing"
)

func TestRun_Success(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "read file output sql",
			args: args{target: "test.sql"},
			want: "SELECT * FROM test;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ReadFile(tt.args.target); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
