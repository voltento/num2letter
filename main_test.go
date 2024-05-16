package main

import (
	"testing"
)

func Test_waysToDecodeNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Single digit valid",
			args: args{s: "1"},
			want: 1,
		},
		{
			name: "Single digit invalid",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "Leading zeros",
			args: args{s: "012"},
			want: 0,
		},
		{
			name: "Double zeros",
			args: args{s: "100"},
			want: 0,
		},
		{
			name: "Example 1",
			args: args{s: "12"},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{s: "226"},
			want: 3,
		},
		{
			name: "Example 3",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{s: "10"},
			want: 1,
		},
		{
			name: "Example 5",
			args: args{s: "2101"},
			want: 1,
		},
		{
			name: "Example 6",
			args: args{s: "1234"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != waysToDecodeNum(tt.args.s) {
				t.Errorf("waysToDecodeNum(%v)", tt.args.s)
			}
		})
	}
}

func Benchmark_waysToDecodeNum(b *testing.B) {
	input := "123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789"
	for i := 0; i < b.N; i++ {
		waysToDecodeNum(input)
	}
}
