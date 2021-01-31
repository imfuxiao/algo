package algo

import (
	"testing"
)

func TestTopologyGraph_Kahn(t *testing.T) {
	//graph := NewTopologyGraph(4)
	//graph.AddEdge(0, 3)
	//graph.AddEdge(3, 2)
	//graph.AddEdge(3, 1)
	//vertexes := graph.Kahn()
	//t.Logf("topology: %+v", vertexes)

	graph := NewTopologyGraph(3)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 1)
	graph.AddEdge(1, 0)
	vertexes := graph.Kahn()
	t.Logf("topology: %+v", vertexes)
}

func TestTopologyGraph_DFS(t *testing.T) {
	graph := NewTopologyGraph(4)
	graph.AddEdge(0, 3)
	graph.AddEdge(3, 2)
	graph.AddEdge(3, 1)
	vertexes := graph.DFS()
	t.Logf("topology: %+v", vertexes)
}
