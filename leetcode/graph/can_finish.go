// 207. 课程表
// 你这个学期必须选修 numCourse 门课程，记为0到numCourse-1 。
//
// 在选修某些课程之前需要一些先修课程。例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package graph

//
func CanFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		inDegree = make([]int, numCourses) // 入度
		queue    = make([]int, 0, numCourses)
	)

	adj := make([][]int, numCourses) // 邻接表: s -> t. 表示s指向t, s先完成
	// init adj, 入度
	for i := range prerequisites {
		p := prerequisites[i]
		adj[p[1]] = append(adj[p[1]], p[0])
		inDegree[p[0]]++
	}

	// 将入度为0的先push到队列中
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		for i, length := 0, len(queue); i < length; i++ {
			// 出队
			// vertex, queue := queue[0], queue[1:] 这样写有bug
			vertex := queue[0]
			queue = queue[1:]
			for j := range adj[vertex] {
				v := adj[vertex][j]
				inDegree[v]--
				if inDegree[v] == 0 {
					queue = append(queue, v)
				}
			}

		}
	}

	for i := range inDegree {
		if inDegree[i] != 0 {
			return false
		}
	}

	return true
}

func CanFinish2(numCourses int, prerequisites [][]int) bool {
	var (
		invertAdj = make([][]int, numCourses) // 逆邻接表: s->t s依赖t, t必须先完成, 即入度
		dfs       func(vertex int)            // 深度优先
		visited   = make([]int, numCourses)   // 记录是否访问: 因为可能存在环, 所以需要多种状态表示: 0: 未开始, 1: 深度搜索中 2: 深度搜索已完成
		result    = true
	)

	// init invertAdj
	for i := range prerequisites {
		invertAdj[prerequisites[i][0]] = append(invertAdj[prerequisites[i][0]], prerequisites[i][1])
	}

	dfs = func(vertex int) {
		visited[vertex] = 1 // 深度搜索进行中
		for i := 0; i < len(invertAdj[vertex]); i++ {
			v := invertAdj[vertex][i]
			if visited[v] == 0 {
				dfs(v)
			} else if visited[v] == 1 { // 表示出现环
				result = false
				break
			}
		}
		visited[vertex] = 2 // 深度搜索已完成
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 { // 0: 表示深度搜索未开始
			dfs(i)
		}
	}

	return result
}
