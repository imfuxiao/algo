package read

import "testing"

func TestExist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: [][]byte{
					{'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a'},
					{'a', 'a', 'a', 'a'},
				},
				word: "aaaaaaaaaaaa",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCEFSADEESE",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCESEEEFS",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkExist-12    	  306595	      3483 ns/op
// BenchmarkExist-12    	  712332	      1621 ns/op
// BenchmarkExist-12    	  530331	      2065 ns/op
// BenchmarkExist-12    	  648775	      1547 ns/op
// BenchmarkExist-12    	 1526186	       778 ns/op
// BenchmarkExist-12    	 3454297	       309 ns/op
func BenchmarkExist(b *testing.B) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCEFSADEESE"
	for i := 0; i < b.N; i++ {
		Exist(board, word)
	}
}

// BenchmarkExist2-12    	 3118449	       340 ns/op
func BenchmarkExist2(b *testing.B) {
}
