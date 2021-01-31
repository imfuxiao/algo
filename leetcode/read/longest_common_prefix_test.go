package read

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			args: args{strs: []string{}},
			want: "",
		},
		{
			args: args{strs: []string{"a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			args: args{strs: []string{}},
			want: "",
		},
		{
			args: args{strs: []string{"a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
