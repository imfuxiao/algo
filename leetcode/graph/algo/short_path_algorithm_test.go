package algo

import (
	"reflect"
	"testing"
)

func TestDirectionWeightGraph_Dijkstra(t *testing.T) {
	graph := NewDirectionWeightGraph(8)
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 3, 1)
	graph.AddEdge(1, 4, 1)
	graph.AddEdge(3, 4, 3)
	if got, want := graph.Dijkstra(0, 4), []int{0, 1, 4}; !reflect.DeepEqual(got, want) {
		t.Fatalf("got:%+v, want:%+v", got, want)
	}

}
