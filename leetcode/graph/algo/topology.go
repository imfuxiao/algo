package algo

type TopologyGraph struct {
	vertexCount int     // 顶点数
	adj         [][]int // 邻接表: s->t 表示t依赖s, s先执行
}

func NewTopologyGraph(vertexCount int) *TopologyGraph {
	return &TopologyGraph{
		vertexCount: vertexCount,
		adj:         make([][]int, vertexCount),
	}
}

// 如果 s 优先 t, 那么就生成一条从s指向t的边edge
func (g *TopologyGraph) AddEdge(s, t int) bool {
	if s < 0 || t < 0 || s >= g.vertexCount || t >= g.vertexCount {
		return false
	}
	g.adj[s] = append(g.adj[s], t)
	return true
}

// kahn算法: 需要用到队列
func (g *TopologyGraph) Kahn() []int {

	var (
		result  []int
		queue   = make([]int, 0, g.vertexCount)
		enqueue = func(vertex int) {
			queue = append(queue, vertex)
		}
		dequeue = func() int {
			vertex := queue[0]
			queue = queue[1:]
			return vertex
		}
	)

	// 1. 统计入度
	inDegree := make([]int, g.vertexCount)
	for i := range g.adj {
		for j := range g.adj[i] {
			// 这里要小心!!!, 注意入度的下标
			inDegree[g.adj[i][j]]++
		}
	}

	// 将入度为0的顶点入队
	for i := range inDegree {
		if inDegree[i] == 0 {
			enqueue(i)
		}
	}

	for len(queue) > 0 {
		// 循环遍历
		for i, length := 0, len(queue); i < length; i++ {
			vertex := dequeue()
			result = append(result, vertex)

			// 获取此顶点指向的顶点, 并将指向顶点的入度-1, 当入度为0, 则添加到队列中
			vertexes := g.adj[vertex]
			for _, v := range vertexes {
				inDegree[v]--
				if inDegree[v] == 0 {
					enqueue(v)
				}
			}

		}
	}

	return result
}

// 深度优先遍历
func (g *TopologyGraph) DFS() []int {
	var (
		// 深度优先递归函数
		dfs func(vertex int)
		// adj 为邻接表: s->t 表示t依赖s, s先执行, 即出度
		// invertAdj 为逆邻接表: s<-t 表示s依赖t, t必须先执行, 即入度
		invertAdj = make([][]int, g.vertexCount) // 逆邻接表: s->t 表示s依赖t, t必须先执行
		visited   = make([]bool, g.vertexCount)  // 访问状态
		result    []int
	)

	// 初始化逆邻接表
	for i := range g.adj {
		for j := range g.adj[i] {
			invertAdj[g.adj[i][j]] = append(invertAdj[g.adj[i][j]], i)
		}
	}

	// 深度搜索遍历递归
	dfs = func(vertex int) {
		for i := 0; i < len(invertAdj[vertex]); i++ {
			if !visited[invertAdj[vertex][i]] {
				visited[invertAdj[vertex][i]] = true
				dfs(invertAdj[vertex][i])
			}
		}
		result = append(result, vertex)
	}

	for i := 0; i < g.vertexCount; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i)
		}
	}

	return result
}
