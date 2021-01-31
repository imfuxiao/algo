package algo

import (
	"math"
)

type Edge struct {
	sid    int // 边的起始顶点编号
	tid    int // 边的终止顶点编号
	weight int // 边的权重
}

func NewEdge(sid, tid, weight int) *Edge {
	return &Edge{
		sid:    sid,
		tid:    tid,
		weight: weight,
	}
}

// Direction weight graph 有向权重图
type DirectionWeightGraph struct {
	vertexCount int       // 顶点的个数
	adj         [][]*Edge // 邻接表
}

func NewDirectionWeightGraph(vertexCount int) *DirectionWeightGraph {
	return &DirectionWeightGraph{
		vertexCount: vertexCount,
		adj:         make([][]*Edge, vertexCount),
	}
}

func (g *DirectionWeightGraph) AddEdge(sid, tid, weight int) {
	g.adj[sid] = append(g.adj[sid], NewEdge(sid, tid, weight))
}

// Dijkstra 最短路径算法
// 需要使用优先队列, 小顶堆
// 此算法与广度优先算法很像, 只是利用优先队列获取最短路径
func (g *DirectionWeightGraph) Dijkstra(sid, tid int) []int {
	var (
		sortPaths = make([]int, g.vertexCount)     // 保存最短路径
		queue     = NewPriorityQueue()             // 优先队列
		visited   = make([]bool, g.vertexCount)    // 记录顶点是否被访问
		vertexes  = make([]*Vertex, g.vertexCount) // 优先队列中的值. 关键变量
	)

	// init vertexes
	for i := range vertexes {
		vertexes[i] = &Vertex{
			id:       i,
			distance: math.MaxInt64,
		}
	}

	for i := range sortPaths {
		sortPaths[i] = -1
	}

	// 源点先入队
	visited[sid] = true
	vertexes[sid].distance = 0
	queue.Push(vertexes[sid])

	for !queue.IsEmpty() {
		previous := queue.Pop()
		if previous.id == tid {
			break
		}
		for i := 0; i < len(g.adj[previous.id]); i++ {
			next := g.adj[previous.id][i]
			distance := previous.distance + next.weight
			if distance < vertexes[next.tid].distance {
				sortPaths[next.tid] = previous.id
				vertexes[next.tid].distance = distance
				if !visited[next.tid] {
					visited[next.tid] = true
					queue.Push(vertexes[next.tid])
				} else {
					queue.Update(vertexes[next.tid])
				}
			}
		}
	}
	return sortPaths
}

type Vertex struct {
	id       int // 顶点ID
	distance int // 起点到此顶点的距离
}

// 优先队列: 小顶堆
type priorityQueue struct {
	vertexes []*Vertex
}

func NewPriorityQueue() *priorityQueue {
	return &priorityQueue{}
}

// 堆化: 根据vertex.distance距离堆化
func (p *priorityQueue) heapfiy() {
	length := len(p.vertexes)
	for i := length>>1 - 1; i >= 0; i-- {
		leftIndex, rightIndex := i<<1+1, i<<1+2
		if leftIndex < length && p.vertexes[leftIndex].distance < p.vertexes[i].distance {
			p.vertexes[i], p.vertexes[leftIndex] = p.vertexes[leftIndex], p.vertexes[i]
		}
		if rightIndex < length && p.vertexes[rightIndex].distance < p.vertexes[i].distance {
			p.vertexes[i], p.vertexes[rightIndex] = p.vertexes[rightIndex], p.vertexes[i]
		}
	}
}

func (p *priorityQueue) Pop() *Vertex {
	v, length := p.vertexes[0], len(p.vertexes)
	p.vertexes[0] = p.vertexes[length-1]
	p.vertexes = p.vertexes[:length-1]
	p.heapfiy()
	return v
}

func (p *priorityQueue) Push(v *Vertex) {
	p.vertexes = append(p.vertexes, v)
	p.heapfiy()
}

func (p *priorityQueue) Update(v *Vertex) {
	length := len(p.vertexes)
	for i := range p.vertexes {
		if p.vertexes[i] == v {
			p.vertexes[i], p.vertexes[length-1] = p.vertexes[length-1], p.vertexes[i]
			break
		}
	}
	p.heapfiy()
}

func (p *priorityQueue) IsEmpty() bool {
	return len(p.vertexes) <= 0
}

func (p *priorityQueue) Empty() {
	p.vertexes = p.vertexes[len(p.vertexes):]
}
