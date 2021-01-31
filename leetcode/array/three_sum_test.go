package array

import (
	"reflect"
	"testing"
)

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{nums: []int{-2, 10, -14, 11, 5, -4, 2, 0, -10, -10, 5, 7, -11, 10, -2, -5, 2, 12, -5, 14, -11, -15, -5, 12, 0, 13, 8, 7, 10, 6, -9, -15, 1, 14, 11, -9, -13, -10, 6, -8, -5, -11, 6, -9, 14, 11, -7, -6, 8, 3, -7, 5, -5, 3, 2, 10, -6, -12, 3, 11, 1, 1, 12, 10, -8, 0, 8, -5, 6, -8, -6, 8, -12, -14, 7, 9, 12, -15, -12, -2, -4, -4, -12, 6, 7, -3, -6, -14, -8, 4, 4, 9, -10, -7, -4, -3, 1, 11, -1, -8, -12, 9, 7, -9, 10, -1, -14, -1, -8, 11, 12, -5, -7}},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			args: args{nums: []int{3, 0, -2, -1, 1, 2}},
			want: [][]int{
				{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1},
			},
		},
		{
			args: args{nums: []int{-2, 0, 1, 1, 2}},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			args: args{nums: []int{}},
			want: nil,
		},
		{
			args: args{nums: []int{0}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_threeSum2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{
		//	args: args{nums: []int{-2, 10, -14, 11, 5, -4, 2, 0, -10, -10, 5, 7, -11, 10, -2, -5, 2, 12, -5, 14, -11, -15, -5, 12, 0, 13, 8, 7, 10, 6, -9, -15, 1, 14, 11, -9, -13, -10, 6, -8, -5, -11, 6, -9, 14, 11, -7, -6, 8, 3, -7, 5, -5, 3, 2, 10, -6, -12, 3, 11, 1, 1, 12, 10, -8, 0, 8, -5, 6, -8, -6, 8, -12, -14, 7, 9, 12, -15, -12, -2, -4, -4, -12, 6, 7, -3, -6, -14, -8, 4, 4, 9, -10, -7, -4, -3, 1, 11, -1, -8, -12, 9, 7, -9, 10, -1, -14, -1, -8, 11, 12, -5, -7}},
		//	want: [][]int{
		//		{-1, -1, 2},
		//		{-1, 0, 1},
		//	},
		//},
		//{
		//	args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
		//	want: [][]int{
		//		{-1, -1, 2},
		//		{-1, 0, 1},
		//	},
		//},
		//{
		//	args: args{nums: []int{3, 0, -2, -1, 1, 2}},
		//	want: [][]int{
		//		{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1},
		//	},
		//},
		//{
		//	args: args{nums: []int{-2, 0, 1, 1, 2}},
		//	want: [][]int{
		//		{-2, 0, 2},
		//		{-2, 1, 1},
		//	},
		//},
		{
			args: args{nums: []int{0, 0, 0, 0}},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			args: args{nums: []int{}},
			want: nil,
		},
		{
			args: args{nums: []int{0}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
