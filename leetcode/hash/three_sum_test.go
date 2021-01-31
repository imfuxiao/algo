package hash

import (
	"reflect"
	"testing"
)

func TestThreeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			},
			want: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			args: args{
				nums: []int{-4, -1, -1, 0, 1, 2},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreeSum2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			},
			want: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			args: args{
				nums: []int{-4, -1, -1, 0, 1, 2},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
