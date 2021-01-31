package graph

import "testing"

func TestCanFinish1(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{2, 0},
					{2, 1},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 7,
				prerequisites: [][]int{
					{1, 0},
					{0, 3},
					{0, 2},
					{3, 2},
					{2, 5},
					{4, 5},
					{5, 6},
					{2, 4},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{0, 1},
					{3, 1},
					{1, 3},
					{3, 2},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{0, 2},
					{2, 1},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 5,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish1(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("CanFinish1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanFinish2(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{2, 0},
					{2, 1},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 7,
				prerequisites: [][]int{
					{1, 0},
					{0, 3},
					{0, 2},
					{3, 2},
					{2, 5},
					{4, 5},
					{5, 6},
					{2, 4},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{0, 1},
					{3, 1},
					{1, 3},
					{3, 2},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 3,
				prerequisites: [][]int{
					{1, 0},
					{0, 2},
					{2, 1},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: true,
		},
		{
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
		{
			args: args{
				numCourses: 5,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish2(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("CanFinish1() = %v, want %v", got, tt.want)
			}
		})
	}
}
