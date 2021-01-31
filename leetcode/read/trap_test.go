package read

import "testing"

func TestTrap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap1(tt.args.height); got != tt.want {
				t.Errorf("Trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrap2(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap2(tt.args.height); got != tt.want {
				t.Errorf("Trap2() = %v, want %v", got, tt.want)
			}
		})
	}
}
