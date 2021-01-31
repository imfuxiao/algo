package binarysearch

import (
	"testing"
)

func TestBinarySearch1(t *testing.T) {
	type args struct {
		n   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	args: args{
		// 		n:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 		num: 5,
		// 	},
		// 	want: 4,
		// },
		{
			args: args{
				n:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				num: 11,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch1(tt.args.n, tt.args.num); got != tt.want {
				t.Errorf("BinarySearch1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch2(t *testing.T) {
	type args struct {
		n   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:   []int{1, 3, 4, 8, 8, 8, 8, 8, 11, 18},
				num: 8,
			},
			want: 3,
		},
		{
			args: args{
				n:   []int{1, 3, 4, 5, 8, 8, 8, 8, 11, 18},
				num: 8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch2(tt.args.n, tt.args.num); got != tt.want {
				t.Errorf("BinarySearch2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch3(t *testing.T) {
	type args struct {
		n   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:   []int{1, 3, 4, 8, 8, 8, 8, 8, 11, 18},
				num: 8,
			},
			want: 7,
		},
		{
			args: args{
				n:   []int{1, 4, 4, 4, 8, 8, 8, 8, 11, 18},
				num: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch3(tt.args.n, tt.args.num); got != tt.want {
				t.Errorf("BinarySearch3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch4(t *testing.T) {
	type args struct {
		n   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:   []int{1, 2, 8, 8, 8, 9, 10, 11, 11, 18},
				num: 6,
			},
			want: 2,
		},
		{
			args: args{
				n:   []int{1, 2, 8, 8, 8, 9, 10, 11, 11, 18},
				num: 12,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch4(tt.args.n, tt.args.num); got != tt.want {
				t.Errorf("BinarySearch4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch5(t *testing.T) {
	type args struct {
		n   []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:   []int{1, 2, 3, 3, 3, 3, 6, 8, 9, 10},
				num: 5,
			},
			want: 5,
		},
		{
			args: args{
				n:   []int{1, 2, 3, 3, 3, 3, 6, 8, 9, 10},
				num: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch5(tt.args.n, tt.args.num); got != tt.want {
				t.Errorf("BinarySearch5() = %v, want %v", got, tt.want)
			}
		})
	}
}
