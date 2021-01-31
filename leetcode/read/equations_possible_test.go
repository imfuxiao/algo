package read

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{
		//	args: args{equations: []string{"a!=b", "b!=c", "c!=a"}},
		//	want: true,
		//},
		//{
		//	args: args{equations: []string{"e==e", "d!=e", "c==d", "d!=e"}},
		//	want: true,
		//},
		//{
		//	args: args{equations: []string{"a==b", "b==a"}},
		//	want: true,
		//},
		{
			args: args{equations: []string{"b==a", "a==b"}},
			want: true,
		},
		//{
		//	args: args{equations: []string{"c==c", "b==d", "x!=z"}},
		//	want: true,
		//},
		//{
		//	args: args{equations: []string{"b==a", "b==d", "x!=z"}},
		//	want: true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
