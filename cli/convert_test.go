package cli

import (
	"os"
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
		stmt *SelectStatement
		want string
	}{
		{
			name: "convert success",
			stmt: &SelectStatement{
				Fields:    []string{"name"},
				TableName: "tbl",
			},
			want: `{"query":{"match_all":{}}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewCli(os.Stdout, os.Stderr).convertToQueryDSL(tt.stmt); got != tt.want {
				t.Errorf("convertESQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
