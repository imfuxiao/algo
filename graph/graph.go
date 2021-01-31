package graph

import (
	"fmt"
)

// 无向图
type Graph struct {
	vertexCount   int     // 顶点个数
	adjacencyList [][]int // 邻接表
}

func NewGraph(vertexCount int) *Graph {
	if vertexCount <= 0 {
		return nil
	}
	return &Graph{
		vertexCount:   vertexCount,
		adjacencyList: make([][]int, vertexCount),
	}
}

// 无向图添加顶点时两个方向都需要添加
func (g *Graph) AddEdge(vertex, otherVertex int) {
	g.adjacencyList[vertex] = append(g.adjacencyList[vertex], otherVertex)
	g.adjacencyList[otherVertex] = append(g.adjacencyList[otherVertex], vertex)
}

func (g *Graph) String() string {
	return fmt.Sprintf("%+v", g.adjacencyList)
}

func (g *Graph) check(start, end int) error {
	if start >= end {
		return fmt.Errorf("start >= end")
	}
	if start < 0 || start >= g.vertexCount {
		return fmt.Errorf("start < 0")
	}

	if end < 0 || end >= g.vertexCount {
		return fmt.Errorf("end < 0")
	}
	return nil
}

// Breadth-First-Search 广度优先搜索算法
func (g *Graph) BFS(start, end int) []int {

	if err := g.check(start, end); err != nil {
		fmt.Printf("%e\n", err)
		return nil
	}

	// 记录顶点vertex 是否已被访问, 初始值false表示未被访问
	visited := make([]bool, g.vertexCount)
	visited[start] = true

	// 记录访问路径, 下标为顶点, 值为父级顶点
	visitPath := make([]int, g.vertexCount)

	// init值为-1,
	for i := range visitPath {
		visitPath[i] = -1
	}

	// 将起始顶点放入队列中
	queue := make(chan int, g.vertexCount)
	queue <- start

	for parentVertex := range queue {
		for j, vertexes := 0, g.adjacencyList[parentVertex]; j < len(vertexes); j++ {
			vertex := vertexes[j]
			if !visited[vertex] {
				visited[vertex] = true
				visitPath[vertex] = parentVertex
				if vertex == end {
					close(queue)
					return getPath(visitPath, start, end)
				}
				queue <- vertex
			}
		}
	}
	return nil
}

func getPath(path []int, start, end int) []int {
	r := make([]int, 0, len(path))
	r = append(r, end)
	for {
		r = append(r, path[end])
		end = path[end]
		if end == start {
			break
		}
	}
	// 因为path中存储的是反向顺序, 所以需要交换位置, 让访问路径从start到end排列
	for i, l := len(r)/2-1, len(r)-1; i >= 0; i-- {
		r[i], r[l-i] = r[l-i], r[i]
	}
	return r
}

// Depth-First-Search 深度优先算法
func (g *Graph) DFS(start, end int) []int {

	if err := g.check(start, end); err != nil {
		fmt.Printf("%e\n", err)
		return nil
	}

	// 记录节点是否已被访问: false 没有
	visited := make([]bool, g.vertexCount)
	visited[start] = true

	// 访问路径
	visitPath := make([]int, g.vertexCount)
	for i := 0; i < g.vertexCount; i++ {
		visitPath[i] = -1
	}

	g.recursionDFS(start, end, visited, visitPath)

	return getPath(visitPath, start, end)
}

func (g *Graph) recursionDFS(start, end int, visited []bool, visitPath []int) {
	for j, vertexes := 0, g.adjacencyList[start]; j < len(vertexes); j++ {
		vertex := vertexes[j]
		if !visited[vertex] {
			visited[vertex] = true
			visitPath[vertex] = start
			if vertex != end {
				g.recursionDFS(vertex, end, visited, visitPath)
			} else {
				return
			}
		}
	}
}
