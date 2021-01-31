package graph

import (
	"reflect"
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	type fields struct {
		g *Graph
	}
	type args struct {
		a     []int
		b     []int
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			fields: fields{
				g: NewGraph(8),
			},
			args: args{
				a:     []int{0, 0, 1, 1, 2, 3, 4, 4, 5, 6},
				b:     []int{1, 3, 2, 4, 5, 4, 5, 6, 7, 7},
				start: 0,
				end:   7,
			},
			want: []int{0, 1, 2, 5, 7},
		},
		{
			fields: fields{
				g: NewGraph(8),
			},
			args: args{
				a:     []int{0, 0, 0, 1, 2, 3, 4, 4, 5, 6},
				b:     []int{1, 3, 4, 2, 5, 5, 5, 7, 6, 7},
				start: 0,
				end:   7,
			},
			want: []int{0, 4, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.a {
				tt.fields.g.AddEdge(v, tt.args.b[i])
			}
			if r := tt.fields.g.BFS(tt.args.start, tt.args.end); !reflect.DeepEqual(r, tt.want) {
				t.Fatalf("you want:%+v, result: %+v", tt.want, r)
			}
		})
	}
}

func TestGraph_DFS(t *testing.T) {
	type fields struct {
		g *Graph
	}
	type args struct {
		a     []int
		b     []int
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			fields: fields{
				g: NewGraph(8),
			},
			args: args{
				a:     []int{0, 0, 1, 1, 2, 3, 4, 4, 5, 6},
				b:     []int{1, 3, 2, 4, 5, 4, 5, 6, 7, 7},
				start: 0,
				end:   7,
			},
			want: []int{0, 1, 2, 5, 4, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.a {
				tt.fields.g.AddEdge(v, tt.args.b[i])
			}
			if got := tt.fields.g.DFS(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
