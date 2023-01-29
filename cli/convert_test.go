package cli

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
			if got, _ := readFile(tt.args.target); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertESQuery_Success(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "convert success",
			args: args{sql: "SELECT * FROM test"},
			want: `{"query": {"match_all": {}}}`,
		},
		{
			name: "convert success",
			args: args{sql: "SEECT * FROM test"},
			want: `{"query": {"match_all": {}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := convertQueryDSL(tt.args.sql); got != tt.want {
				t.Errorf("convertESQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
