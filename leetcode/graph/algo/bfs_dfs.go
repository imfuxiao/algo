package algo

// 无向图
type Graph struct {
	v   int     // 顶点个数
	adj [][]int // 邻接表: 下标为顶点(用数字表示), 下标对应的值是数组, 数组中每个值对应与这个顶点关联的点
}

func NewGraph(v int) *Graph {
	return &Graph{
		v:   v,
		adj: make([][]int, v),
	}
}

// vertexFirst, vertexSecond: 无向图的两个点
// 无向图一个边需要添加两次
func (g *Graph) AddEdge(vertexFirst, vertexSecond int) bool {
	if vertexFirst >= g.v || vertexSecond >= g.v {
		return false
	}
	g.adj[vertexFirst] = append(g.adj[vertexFirst], vertexSecond)
	g.adj[vertexSecond] = append(g.adj[vertexSecond], vertexFirst)
	return true
}

// 广度优先算法
// src: 源顶点
// dist: 目标顶点
// 返回从src到dist的路径
func (g *Graph) BFS(src, dist int) ([]int, bool) {
	if src >= g.v || dist >= g.v {
		return nil, false
	}
	var (
		paths        []int
		queue        = make([]int, 0, g.v) // 队列保存每一层的顶点
		visited      = make([]bool, g.v)   // 标识顶点是否已被访问
		visitPath    = make([]int, g.v)    // 记录访问路径
		getVisitPath func([]int) []int
	)

	getVisitPath = func(visitPath []int) []int {
		result := []int{dist}
		d := dist
		for {
			result = append(result, visitPath[d])
			if visitPath[d] == src {
				break
			}
			d = visitPath[d]
		}
		for i, length := 0, len(result); i < length/2; i++ {
			result[i], result[length-i-1] = result[length-i-1], result[i]
		}
		return result
	}

	enqueue := func(vertex int) {
		queue = append(queue, vertex)
	}
	dequeue := func() int {
		vertex := -1
		if len(queue) > 0 {
			vertex = queue[0]
			queue = queue[1:]
		}
		return vertex
	}

	for i := range visitPath {
		visitPath[i] = -1
	}

	visited[src] = true
	enqueue(src) // 入队访问节点

	for len(queue) > 0 {
		for length := len(queue); length > 0; length-- {
			vertex := dequeue()
			dists := g.adj[vertex]
			for i := 0; i < len(dists); i++ {
				if visited[dists[i]] {
					continue
				}
				visitPath[dists[i]] = vertex
				if dists[i] == dist {
					paths = getVisitPath(visitPath)
					return paths, true
				}
				visited[dists[i]] = true
				enqueue(dists[i])
			}
		}
	}
	return paths, false
}

// 深度优先算法
// src: 源顶点
// dist: 目标顶点
func (g *Graph) DFS(src, dist int) ([]int, bool) {
	if src >= g.v || dist >= g.v || src == dist {
		return nil, false
	}

	var (
		visited     = make([]bool, g.v) // 标识当前顶点是否已经被访问
		visitedPath = make([]int, g.v)  // 访问路径
		dfs         func(src, dist int)
		getPath     func(visitedPath []int)
		path        = make([]int, 0, g.v)
	)

	for i := range visitedPath {
		visitedPath[i] = -1
	}
	visited[src] = true

	getPath = func(visitedPath []int) {
		path = append(path, dist)
		d := dist
		for {
			path = append(path, visitedPath[d])
			if visitedPath[d] == src {
				break
			}
			d = visitedPath[d]
		}
	}
	isFound := false
	dfs = func(src, dist int) {
		if src == dist {
			getPath(visitedPath)
			isFound = true
			return
		}
		visited[src] = true
		for _, v := range g.adj[src] {
			if !visited[v] {
				visited[v] = true
				visitedPath[v] = src
				dfs(v, dist)
			}
			if isFound {
				break
			}
		}
	}

	dfs(src, dist)

	return path, len(path) != 0
}
