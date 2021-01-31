package algo

import (
	"testing"
)

func TestGraph_BFS(t *testing.T) {
	bfsGraph := NewGraph(8)
	bfsGraph.AddEdge(0, 1)
	bfsGraph.AddEdge(0, 3)

	bfsGraph.AddEdge(1, 2)
	bfsGraph.AddEdge(1, 4)

	bfsGraph.AddEdge(2, 5)

	bfsGraph.AddEdge(3, 4)

	bfsGraph.AddEdge(4, 5)
	bfsGraph.AddEdge(4, 6)

	bfsGraph.AddEdge(5, 7)

	bfsGraph.AddEdge(6, 7)

	if path, notReach := bfsGraph.BFS(0, 7); !notReach {
		t.Fatalf("src 0, dist 7 should reach")
	} else {
		t.Logf("visited path: %+v", path)
	}

	if path, notReach := bfsGraph.BFS(0, 0); notReach {
		t.Fatalf("src 0, dist 7 should reach")
	} else {
		t.Logf("visited path: %+v", path)
	}
}

func TestGraph_DFS(t *testing.T) {
	bfsGraph := NewGraph(8)
	bfsGraph.AddEdge(0, 1)
	bfsGraph.AddEdge(0, 3)

	bfsGraph.AddEdge(1, 2)
	bfsGraph.AddEdge(1, 4)

	bfsGraph.AddEdge(2, 5)

	bfsGraph.AddEdge(3, 4)

	bfsGraph.AddEdge(4, 5)
	bfsGraph.AddEdge(4, 6)

	bfsGraph.AddEdge(5, 7)

	bfsGraph.AddEdge(6, 7)

	if path, notReach := bfsGraph.DFS(0, 7); !notReach {
		t.Fatalf("src 0, dist 7 should reach")
	} else {
		t.Logf("visited path: %+v", path)
	}

	if path, notReach := bfsGraph.DFS(0, 0); notReach {
		t.Fatalf("src 0, dist 7 should reach")
	} else {
		t.Logf("visited path: %+v", path)
	}
}
