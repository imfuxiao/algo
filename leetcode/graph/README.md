## 图

知识要求:

* 实现有向图, 无向图, 有权图, 无权图的邻接矩阵和邻接表示法.
* 实现图的深度优先搜索, 广度优先搜索
* 实现Dijkstra算法, A*算法
* 实现拓扑排序的Kahn算法. DFS算法

## 知识详解

相关课程: 第30课, 第31课, 第43课, 第44课, 第49课.

图(Graph), 顶点(vertex), 边(edge), 度(degree).

无向图

有向图: 出度, 入度

带权图: 权重

图的存储方法: 

1. 邻接矩阵(Adjacency Matrix): 二维数组, 劣势: 浪费空间

2. 邻接表(Adjacency List): 1个顶点对应一个链表.

搜索算法: 

1. 广度优先搜索算法, Breadth First Search, BFS. 需要使用队列实现.
2. 深度优先搜索算法, Depth First Search, DFS. 回溯+递归.

拓扑排序: 

1. 凡是需要通过局部顺序来推导全局顺序的, 一般都能用拓扑排序解决.
2. 拓扑排序还能检测图中环的存在.
3. 拓扑排序的两种算法: Kahn, DFS

最短路径算法(Shortest Path Algorithm):

1. 建模: 将复杂的场景抽象成具体的数据结构;
    
    例如, 地图可以抽象为一个有向权图. 每个岔路口为一个顶点. 距离为权重, 顶点之间双行线则为双向边, 否则为单向边.
   
    则最短路径问题转化为: 在一个有向权图中, 求两个顶点之间的最短路径.

2. 算法: Dijkstra算法. 单源路径算法(一个顶点到另一个顶点)

   此算法有点类似BFS(广度优先算法), 只是把普通队列换成了一个**优先队列**, 通过优先队列获取每次更新后的最短路径.

3. A*(AStart)算法: A*算法是对Dijkstra算法的改造




## LeetCode习题

* 200. [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/description/)
    
难度: 中等

* 36. [有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)
    
难度: 中等

* 207. [课程表](https://leetcode-cn.com/problems/course-schedule/)
    
难度: 中等