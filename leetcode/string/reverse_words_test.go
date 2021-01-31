package string

import (
	"reflect"
	"testing"
)

func TestReverseWords1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "the sky is blue"},
			want: "blue is sky the",
		},
		{
			args: args{s: " the   sky is blue  "},
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords1(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseWords1()=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestReverseWords2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	args: args{s: "the sky is blue"},
		//	want: "blue is sky the",
		//},
		{
			args: args{s: " the   sky is blue  "},
			want: "blue is sky the",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords2(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords2() = %v, want %v", got, tt.want)
			}
		})
	}
}
